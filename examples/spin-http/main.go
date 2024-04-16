package main

import (
	"net/http"

	spinHttp "github.com/fermyon/spin-go-sdk/pkg/http"
)

func init() {
	spinHttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("Hello, world")); err != nil {
			panic(err.Error())
		}

		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(200)
	})
}

func main() {}
