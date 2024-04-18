// Code generated by wit-bindgen-go. DO NOT EDIT.

//go:build !wasip1

// Package terminalstdout represents the interface "wasi:cli/terminal-stdout@0.2.0-rc-2023-11-10".
//
// An interface providing an optional `terminal-output` for stdout as a
// link-time authority.
package terminalstdout

import (
	terminaloutput "github.com/fermyon/spin-go-sdk/internal/wasi/cli/v0.2.0-rc-2023-11-10/terminal-output"
	"github.com/ydnar/wasm-tools-go/cm"
)

// TerminalOutput represents the resource "wasi:cli/terminal-output@0.2.0-rc-2023-11-10#terminal-output".
//
// See [terminaloutput.TerminalOutput] for more information.
type TerminalOutput = terminaloutput.TerminalOutput

// GetTerminalStdout represents function "wasi:cli/terminal-stdout@0.2.0-rc-2023-11-10#get-terminal-stdout".
//
// If stdout is connected to a terminal, return a `terminal-output` handle
// allowing further interaction with it.
//
//	get-terminal-stdout: func() -> option<terminal-output>
//
//go:nosplit
func GetTerminalStdout() cm.Option[TerminalOutput] {
	var result cm.Option[TerminalOutput]
	wasmimport_GetTerminalStdout(&result)
	return result
}

//go:wasmimport wasi:cli/terminal-stdout@0.2.0-rc-2023-11-10 get-terminal-stdout
//go:noescape
func wasmimport_GetTerminalStdout(result *cm.Option[TerminalOutput])

type Interface interface {
	GetTerminalStdout() cm.Option[TerminalOutput]
}

var instance Interface
