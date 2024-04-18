// Code generated by wit-bindgen-go. DO NOT EDIT.

//go:build !wasip1

// Package terminalinput represents the interface "wasi:cli/terminal-input@0.2.0-rc-2023-11-10".
package terminalinput

import (
	"github.com/ydnar/wasm-tools-go/cm"
)

// TerminalInput represents the resource "wasi:cli/terminal-input@0.2.0-rc-2023-11-10#terminal-input".
//
// The input side of a terminal.
//
//	resource terminal-input
type TerminalInput cm.Resource

// ResourceDrop represents the Canonical ABI function "resource-drop".
//
// Drops a resource handle.
//
//go:nosplit
func (self TerminalInput) ResourceDrop() {
	self.wasmimport_ResourceDrop()
}

//go:wasmimport wasi:cli/terminal-input@0.2.0-rc-2023-11-10 [resource-drop]terminal-input
//go:noescape
func (self TerminalInput) wasmimport_ResourceDrop()

type Interface interface {
}

var instance Interface
