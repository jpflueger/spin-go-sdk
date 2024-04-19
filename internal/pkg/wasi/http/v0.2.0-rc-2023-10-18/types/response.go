package types

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/fermyon/spin-go-sdk/internal/pkg/wasi/io/v0.2.0-rc-2023-10-18/streams"
	"github.com/fermyon/spin-go-sdk/internal/wasi/http/v0.2.0-rc-2023-10-18/types"
	"github.com/ydnar/wasm-tools-go/cm"
)

type ResponseOutparam = types.ResponseOutparam

// convert the ResponseOutparam to http.ResponseWriter
func NewHttpResponseWriter(outparam ResponseOutparam) *ResponseOutparamWriter {
	return &ResponseOutparamWriter{
		outparam: outparam,
	}
}

// A response writer keeps track of response state until it is ready to be "flushed" to the host after handling
type ResponseOutparamWriter struct {
	// wasi response outparam is set at the end of http_trigger_handle
	outparam ResponseOutparam
	// buffered response body because we can only write once we know the http-status and headers
	buffer bytes.Buffer
	// response status code
	statusCode int
	// go headers are reconciled on call to WriteHeader, Flush or at the end of http_trigger_handle
	headers http.Header
}

// implement net/http/ResponseWriter.Header()
func (w *ResponseOutparamWriter) Header() http.Header {
	return w.headers
}

// implement net/http/ResponseWriter.WriteHeader(statusCode int)
func (w *ResponseOutparamWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
}

// implement net/http/ResponseWriter.Write(buf []byte)
func (w *ResponseOutparamWriter) Write(buf []byte) (int, error) {
	return w.buffer.Write(buf)
}

// implement net/http/Flusher.Flush()
func (w *ResponseOutparamWriter) Flush() {
	var headers types.Headers
	CopyToWasiFields(w.headers, &headers)

	response := types.NewOutgoingResponse(types.StatusCode(w.statusCode), headers)

	var result cm.ErrResult[types.OutgoingResponse, types.Error]
	if bodyResult := response.Write(); bodyResult.IsErr() {
		msg := fmt.Sprintf("failed to acquire outgoing-body resource for writing: %s\n", bodyResult.Err())
		result = cm.Err[cm.ErrResult[types.OutgoingResponse, types.Error]](types.ErrorUnexpectedError(msg))
	} else if streamResult := bodyResult.OK().Write(); streamResult.IsErr() {
		msg := fmt.Sprintf("failed to acquire outgoing-body resource for writing: %s\n", streamResult.Err())
		result = cm.Err[cm.ErrResult[types.OutgoingResponse, types.Error]](types.ErrorUnexpectedError(msg))
	} else {
		s := streamResult.OK()
		sw := streams.NewOutputStreamWriter(*s)
		if _, err := w.buffer.WriteTo(sw); err != nil {
			msg := fmt.Sprintf("failed to write response buffer to stream: %s\n", err.Error())
			result = cm.Err[cm.ErrResult[types.OutgoingResponse, types.Error]](types.ErrorUnexpectedError(msg))
		} else if flushResult := s.Flush(); flushResult.IsErr() {
			flushErr := flushResult.Err()
			if flushErr.Closed() {
				msg := fmt.Sprintln("failed to flush a closed response stream")
				result = cm.Err[cm.ErrResult[types.OutgoingResponse, types.Error]](types.ErrorUnexpectedError(msg))
			} else {
				msg := fmt.Sprintf("failed to flush response stream, last operation failed: %s\n", flushErr.LastOperationFailed().ToDebugString())
				result = cm.Err[cm.ErrResult[types.OutgoingResponse, types.Error]](types.ErrorUnexpectedError(msg))
			}
		} else {
			result = cm.OK[cm.ErrResult[types.OutgoingResponse, types.Error]](response)
		}
	}

	types.ResponseOutparamSet(w.outparam, result)
}
