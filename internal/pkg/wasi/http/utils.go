package http

import (
	"fmt"
	"net/http"

	wasi "github.com/fermyon/spin-go-sdk/internal/wasi/http/types"
	"github.com/ydnar/wasm-tools-go/cm"
)

func methodToString(m wasi.Method) (string, error) {
	if m.Connect() {
		return "CONNECT", nil
	}
	if m.Delete() {
		return "DELETE", nil
	}
	if m.Get() {
		return "GET", nil
	}
	if m.Head() {
		return "HEAD", nil
	}
	if m.Options() {
		return "OPTIONS", nil
	}
	if m.Patch() {
		return "PATCH", nil
	}
	if m.Post() {
		return "POST", nil
	}
	if m.Put() {
		return "PUT", nil
	}
	if m.Trace() {
		return "TRACE", nil
	}
	if other := m.Other(); other != nil {
		return *other, fmt.Errorf("unknown http method 'other'")
	}
	return "", fmt.Errorf("failed to convert http method")
}

func toHttpHeader(src wasi.Headers, dest *http.Header) {
	for _, elem := range src.Entries().Slice() {
		key := string(elem.F0)
		val := string(cm.List[uint8](elem.F1).Slice())
		dest.Add(key, val)
	}
}
