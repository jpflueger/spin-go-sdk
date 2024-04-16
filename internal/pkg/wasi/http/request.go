package http

import (
	"fmt"
	"io"
	"net/http"

	streams "github.com/fermyon/spin-go-sdk/internal/pkg/wasi/io/streams"
	wasi "github.com/fermyon/spin-go-sdk/internal/wasi/http/types"
)

// convert the IncomingRequest to http.Request
func NewHttpRequest(ir wasi.IncomingRequest) (req *http.Request, err error) {
	// convert the http method to string
	method, err := methodToString(ir.Method())
	if err != nil {
		return nil, err
	}

	// convert the path with query to a url
	var url string
	if pathWithQuery := ir.PathWithQuery(); pathWithQuery.None() {
		url = ""
	} else {
		url = *pathWithQuery.Some()
	}

	// convert the body to a reader
	var body io.Reader
	if consumeResult := ir.Consume(); consumeResult.IsErr() {
		return nil, fmt.Errorf("failed to consume incoming request %s", *consumeResult.Err())
	} else if streamResult := consumeResult.OK().Stream(); streamResult.IsErr() {
		return nil, fmt.Errorf("failed to consume incoming requests's stream %s", streamResult.Err())
	} else {
		body = streams.NewReader(*streamResult.OK())
	}

	// create a new request
	req, err = http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	// update additional fields
	toHttpHeader(ir.Headers(), &req.Header)

	return req, nil
}
