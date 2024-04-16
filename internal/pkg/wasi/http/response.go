package http

import (
	"fmt"
	"net/http"
	"strings"

	wasi "github.com/fermyon/spin-go-sdk/internal/wasi/http/types"
	"github.com/ydnar/wasm-tools-go/cm"
)

type responseOutparamWriter struct {
	// wasi response outparam is set at the end of http_trigger_handle
	outparam wasi.ResponseOutparam
	// wasi response
	response wasi.OutgoingResponse
	// wasi http headers
	wasiHeaders wasi.Headers
	// go httpHeaders are reconciled on call to WriteHeader, Flush or at the end of http_trigger_handle
	httpHeaders http.Header
	// wasi response body is set on first write because it can only be called once
	body *wasi.OutgoingBody
	// wasi response stream is set on first write because it can only be called once
	stream *wasi.OutputStream
}

func (row responseOutparamWriter) Header() http.Header {
	return row.httpHeaders
}

func (row responseOutparamWriter) Write(buf []byte) (int, error) {
	// acquire the response body's resource handle on first call to write
	if row.body == nil {
		if bodyResult := row.response.Body(); bodyResult.IsErr() {
			return 0, fmt.Errorf("failed to acquire resource handle to response body: %s", bodyResult.Err())
		} else {
			row.body = bodyResult.OK()
		}

		if writeResult := row.body.Write(); writeResult.IsErr() {
			return 0, fmt.Errorf("failed to acquire resource handle for response body's stream: %s", writeResult.Err())
		} else {
			row.stream = writeResult.OK()
		}
	}

	//TODO: determine if we need to do these to fulfill the ResponseWriter contract
	// call WriteHeader(http.StatusOK) if it hasn't been called yet
	// call DetectContentType if headers doesn't contain content-type yet
	// if total data is under "a few" KB and there are no flush calls, Content-Length is added automatically

	contents := cm.ToList(buf)
	if writeResult := row.stream.Write(contents); writeResult.IsErr() {
		if writeResultErr := writeResult.Err(); writeResultErr.Closed() {
			return 0, fmt.Errorf("failed to write to response body's stream: closed")
		} else {
			//TODO: possible nil error here
			return 0, fmt.Errorf("failed to write to response body's stream: %s", writeResultErr.LastOperationFailed().ToDebugString())
		}
	}
	return int(contents.Len()), nil
}

func (row responseOutparamWriter) WriteHeader(statusCode int) {
	row.response.SetStatusCode(wasi.StatusCode(statusCode))
	if err := row.reconcileHeaders(); err != nil {
		panic(err.Error())
	}
}

// reconcile headers from go to wasi
func (row responseOutparamWriter) reconcileHeaders() error {
	for key, vals := range row.httpHeaders {
		// convert each value distincly
		fieldVals := make([]wasi.FieldValue, 0, len(vals))
		for i, val := range vals {
			fieldVals[i] = wasi.FieldValue(cm.ToList([]uint8(val)))
		}

		if result := row.wasiHeaders.Set(wasi.FieldKey(key), cm.ToList(fieldVals)); result.IsErr() {
			switch *result.Err() {
			case wasi.HeaderErrorInvalidSyntax:
				return fmt.Errorf("failed to set header %s to [%s]: invalid syntax", key, strings.Join(vals, ","))
			case wasi.HeaderErrorForbidden:
				return fmt.Errorf("failed to set forbidden header key %s", key)
			case wasi.HeaderErrorImmutable:
				return fmt.Errorf("failed to set header on immutable header fields")
			}
		}
	}

	//TODO: handle deleted headers

	return nil
}

// convert the ResponseOutparam to http.ResponseWriter
func NewHttpResponseWriter(out wasi.ResponseOutparam) http.ResponseWriter {
	// need to keep the headers because this is mutable
	headers := wasi.NewFields()
	return responseOutparamWriter{
		outparam:    out,
		response:    wasi.NewOutgoingResponse(headers),
		wasiHeaders: headers,
		httpHeaders: http.Header{},
	}
}
