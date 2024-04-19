package streams

import (
	"fmt"

	"github.com/fermyon/spin-go-sdk/internal/wasi/http/v0.2.0-rc-2023-10-18/types"
	"github.com/ydnar/wasm-tools-go/cm"
)

type OutputStream = types.OutputStream

type OutputStreamWriter struct {
	stream types.OutputStream
}

func NewOutputStreamWriter(s types.OutputStream) *OutputStreamWriter {
	return &OutputStreamWriter{
		stream: s,
	}
}

// implement io.Writer
func (w *OutputStreamWriter) Write(p []byte) (n int, err error) {
	contents := cm.ToList([]uint8(p))
	if result := w.stream.Write(contents); result.IsErr() {
		err := result.Err()
		if err.Closed() {
			return 0, fmt.Errorf("failed to write to closed stream")
		} else {
			return 0, fmt.Errorf("failed to write to stream, last operation failed: %s", err.LastOperationFailed().ToDebugString())
		}
	}
	return int(contents.Len()), nil
}
