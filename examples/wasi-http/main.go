package main

import (
	"fmt"

	"github.com/fermyon/spin-go-sdk/generated/wasi/http/types"
)

//export wasi:http/incoming-handler#handle
func handle(_ types.IncomingRequest, _ types.ResponseOutparam) {
	fmt.Println("Hello, world!")
}

func main() {
}
