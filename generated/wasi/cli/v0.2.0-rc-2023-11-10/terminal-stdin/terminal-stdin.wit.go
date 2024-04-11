// Code generated by wit-bindgen-go. DO NOT EDIT.

//go:build !wasip1

// Package terminalstdin represents the interface "wasi:cli/terminal-stdin@0.2.0-rc-2023-11-10".
//
// An interface providing an optional `terminal-input` for stdin as a
// link-time authority.
package terminalstdin

import (
	terminalinput "github.com/fermyon/spin-go-sdk/generated/wasi/cli/v0.2.0-rc-2023-11-10/terminal-input"
	"github.com/ydnar/wasm-tools-go/cm"
)

// TerminalInput represents the resource "wasi:cli/terminal-input@0.2.0-rc-2023-11-10#terminal-input".
//
// See [terminalinput.TerminalInput] for more information.
type TerminalInput = terminalinput.TerminalInput

// GetTerminalStdin represents function "wasi:cli/terminal-stdin@0.2.0-rc-2023-11-10#get-terminal-stdin".
//
// If stdin is connected to a terminal, return a `terminal-input` handle
// allowing further interaction with it.
//
//	get-terminal-stdin: func() -> option<terminal-input>
//
//go:nosplit
func GetTerminalStdin() cm.Option[TerminalInput] {
	var result cm.Option[TerminalInput]
	wasmimport_GetTerminalStdin(&result)
	return result
}

//go:wasmimport wasi:cli/terminal-stdin@0.2.0-rc-2023-11-10 get-terminal-stdin
//go:noescape
func wasmimport_GetTerminalStdin(result *cm.Option[TerminalInput])

type Interface interface {
	GetTerminalStdin() cm.Option[TerminalInput]
}

var instance Interface
