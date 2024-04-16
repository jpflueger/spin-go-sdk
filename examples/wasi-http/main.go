package main

import (
	"fmt"
	"net/http"

	spinHttp "github.com/fermyon/spin-go-sdk/pkg/http"
)

func init() {
	spinHttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello, world")
	})
}

func main() {}
