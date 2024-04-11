package main

import (
	"fmt"

	wasihttp "github.com/fermyon/spin-go-sdk/generated/wasi/http/v0.2.0/types"
)

//export wasi:http/incoming-handler@0.2.0#handle
func handle(_ wasihttp.IncomingRequest, _ wasihttp.ResponseOutparam) {
	fmt.Println("Something happened!")
}

func main() {}
