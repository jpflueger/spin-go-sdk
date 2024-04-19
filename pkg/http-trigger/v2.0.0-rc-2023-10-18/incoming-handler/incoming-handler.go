package http

import (
	"fmt"
	"net/http"
	"os"

	"github.com/fermyon/spin-go-sdk/internal/pkg/wasi/http/v0.2.0-rc-2023-10-18/types"
)

// defaultHandler is a placeholder for returning a useful error to stderr when
// the handler is not set.
var defaultHandler = func(http.ResponseWriter, *http.Request) {
	fmt.Fprintln(os.Stderr, "http handler undefined")
}

// handler is the function that will be called by the http trigger in Spin.
var handler = defaultHandler

// Handle sets the handler function for the http trigger.
// It must be set in an init() function.
func Handle(fn func(http.ResponseWriter, *http.Request)) {
	handler = fn
}

//lint:ignore U1000 Ignore unused function for wasi export
//go:export wasi:http/incoming-handler@0.2.0-rc-2023-10-18#handle
//go:wasmexport wasi:http/incoming-handler@0.2.0-rc-2023-10-18#handle
func http_trigger_handle(req types.IncomingRequest, res types.ResponseOutparam) {
	// convert the incoming request to go's net/http type
	httpReq, err := types.NewHttpRequest(req)
	if err != nil {
		fmt.Printf("failed to convert wasi/http/types.IncomingRequest to http.Request: %s\n", err)
		return
	}

	// convert the response outparam to go's net/http type
	httpResponseWriter := types.NewHttpResponseWriter(res)

	// run the user's handler
	handler(httpResponseWriter, httpReq)

	// flush the response to the host
	httpResponseWriter.Flush()
}
