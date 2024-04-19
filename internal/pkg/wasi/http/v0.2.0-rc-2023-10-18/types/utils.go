package types

import (
	"fmt"
	"net/http"

	"github.com/fermyon/spin-go-sdk/internal/wasi/http/v0.2.0-rc-2023-10-18/types"
	"github.com/ydnar/wasm-tools-go/cm"
)

// re-export types to make it easier to use from the handler package
type Method = types.Method
type Headers = types.Headers

func MethodToString(m Method) (string, error) {
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

func CopyToHttpHeader(src Headers, dest *http.Header) {
	for _, elem := range src.Entries().Slice() {
		key := string(elem.F0)
		val := string(cm.List[uint8](elem.F1).Slice())
		dest.Add(key, val)
	}
}

func CopyToWasiFields(src http.Header, dest *types.Headers) {
	for key, vals := range src {
		// convert each value distincly
		fieldVals := make([]cm.List[uint8], 0, len(vals))
		for i, val := range vals {
			fieldVals[i] = cm.ToList([]uint8(val))
		}

		dest.Set(key, cm.ToList(fieldVals))
	}
}
