// Code generated by wit-bindgen-go. DO NOT EDIT.

//go:build !wasip1

// Package types represents the interface "wasi:http/types@0.2.0-rc-2023-10-18".
//
// The `wasi:http/types` interface is meant to be imported by components to
// define the HTTP resource types and operations used by the component's
// imported and exported interfaces.
package types

import (
	"github.com/fermyon/spin-go-sdk/internal/wasi/io/v0.2.0-rc-2023-10-18/poll"
	"github.com/fermyon/spin-go-sdk/internal/wasi/io/v0.2.0-rc-2023-10-18/streams"
	"github.com/ydnar/wasm-tools-go/cm"
)

// Error represents the variant "wasi:http/types@0.2.0-rc-2023-10-18#error".
//
// TODO: perhaps better align with HTTP semantics?
// This type enumerates the different kinds of errors that may occur when
// initially returning a response.
//
//	variant error {
//		invalid-url(string),
//		timeout-error(string),
//		protocol-error(string),
//		unexpected-error(string),
//	}
type Error cm.Variant[uint8, string, string]

// ErrorInvalidURL returns a [Error] of case "invalid-url".
func ErrorInvalidURL(data string) Error {
	return cm.New[Error](0, data)
}

// InvalidURL returns a non-nil *[string] if [Error] represents the variant case "invalid-url".
func (self *Error) InvalidURL() *string {
	return cm.Case[string](self, 0)
}

// ErrorTimeoutError returns a [Error] of case "timeout-error".
func ErrorTimeoutError(data string) Error {
	return cm.New[Error](1, data)
}

// TimeoutError returns a non-nil *[string] if [Error] represents the variant case "timeout-error".
func (self *Error) TimeoutError() *string {
	return cm.Case[string](self, 1)
}

// ErrorProtocolError returns a [Error] of case "protocol-error".
func ErrorProtocolError(data string) Error {
	return cm.New[Error](2, data)
}

// ProtocolError returns a non-nil *[string] if [Error] represents the variant case "protocol-error".
func (self *Error) ProtocolError() *string {
	return cm.Case[string](self, 2)
}

// ErrorUnexpectedError returns a [Error] of case "unexpected-error".
func ErrorUnexpectedError(data string) Error {
	return cm.New[Error](3, data)
}

// UnexpectedError returns a non-nil *[string] if [Error] represents the variant case "unexpected-error".
func (self *Error) UnexpectedError() *string {
	return cm.Case[string](self, 3)
}

// Fields represents the resource "wasi:http/types@0.2.0-rc-2023-10-18#fields".
//
// This following block defines the `fields` resource which corresponds to
// HTTP standard Fields.
//
//	resource fields
type Fields cm.Resource

// ResourceDrop represents the Canonical ABI function "resource-drop".
//
// Drops a resource handle.
//
//go:nosplit
func (self Fields) ResourceDrop() {
	self.wasmimport_ResourceDrop()
}

//go:wasmimport wasi:http/types@0.2.0-rc-2023-10-18 [resource-drop]fields
//go:noescape
func (self Fields) wasmimport_ResourceDrop()

// NewFields represents constructor "[constructor]fields".
//
// Multiple values for a header are multiple entries in the list with the
// same key.
//
//	[constructor]fields(entries: list<tuple<string, list<u8>>>)
//
//go:nosplit
func NewFields(entries cm.List[cm.Tuple[string, cm.List[uint8]]]) Fields {
	return wasmimport_NewFields(entries)
}

//go:wasmimport wasi:http/types@0.2.0-rc-2023-10-18 [constructor]fields
//go:noescape
func wasmimport_NewFields(entries cm.List[cm.Tuple[string, cm.List[uint8]]]) Fields

// Append represents method "append".
//
//	append: func(name: string, value: list<u8>)
//
//go:nosplit
func (self Fields) Append(name string, value cm.List[uint8]) {
	self.wasmimport_Append(name, value)
}

//go:wasmimport wasi:http/types@0.2.0-rc-2023-10-18 [method]fields.append
//go:noescape
func (self Fields) wasmimport_Append(name string, value cm.List[uint8])

// Clone represents method "clone".
//
// Deep copy of all contents in a fields.
//
//	clone: func() -> fields
//
//go:nosplit
func (self Fields) Clone() Fields {
	return self.wasmimport_Clone()
}

//go:wasmimport wasi:http/types@0.2.0-rc-2023-10-18 [method]fields.clone
//go:noescape
func (self Fields) wasmimport_Clone() Fields

// Delete represents method "delete".
//
//	delete: func(name: string)
//
//go:nosplit
func (self Fields) Delete(name string) {
	self.wasmimport_Delete(name)
}

//go:wasmimport wasi:http/types@0.2.0-rc-2023-10-18 [method]fields.delete
//go:noescape
func (self Fields) wasmimport_Delete(name string)

// Entries represents method "entries".
//
// Values off wire are not necessarily well formed, so they are given by
// list<u8> instead of string.
//
//	entries: func() -> list<tuple<string, list<u8>>>
//
//go:nosplit
func (self Fields) Entries() cm.List[cm.Tuple[string, cm.List[uint8]]] {
	var result cm.List[cm.Tuple[string, cm.List[uint8]]]
	self.wasmimport_Entries(&result)
	return result
}

//go:wasmimport wasi:http/types@0.2.0-rc-2023-10-18 [method]fields.entries
//go:noescape
func (self Fields) wasmimport_Entries(result *cm.List[cm.Tuple[string, cm.List[uint8]]])

// Get represents method "get".
//
// Values off wire are not necessarily well formed, so they are given by
// list<u8> instead of string.
//
//	get: func(name: string) -> list<list<u8>>
//
//go:nosplit
func (self Fields) Get(name string) cm.List[cm.List[uint8]] {
	var result cm.List[cm.List[uint8]]
	self.wasmimport_Get(name, &result)
	return result
}

//go:wasmimport wasi:http/types@0.2.0-rc-2023-10-18 [method]fields.get
//go:noescape
func (self Fields) wasmimport_Get(name string, result *cm.List[cm.List[uint8]])

// Set represents method "set".
//
// Values off wire are not necessarily well formed, so they are given by
// list<u8> instead of string.
//
//	set: func(name: string, value: list<list<u8>>)
//
//go:nosplit
func (self Fields) Set(name string, value cm.List[cm.List[uint8]]) {
	self.wasmimport_Set(name, value)
}

//go:wasmimport wasi:http/types@0.2.0-rc-2023-10-18 [method]fields.set
//go:noescape
func (self Fields) wasmimport_Set(name string, value cm.List[cm.List[uint8]])

// FutureIncomingResponse represents the resource "wasi:http/types@0.2.0-rc-2023-10-18#future-incoming-response".
//
// The following block defines a special resource type used by the
// `wasi:http/outgoing-handler` interface to emulate
// `future<result<response, error>>` in advance of Preview3. Given a
// `future-incoming-response`, the client can call the non-blocking `get`
// method to get the result if it is available. If the result is not available,
// the client can call `listen` to get a `pollable` that can be passed to
// `wasi:io/poll.poll-list`.
//
//	resource future-incoming-response
type FutureIncomingResponse cm.Resource

// ResourceDrop represents the Canonical ABI function "resource-drop".
//
// Drops a resource handle.
//
//go:nosplit
func (self FutureIncomingResponse) ResourceDrop() {
	self.wasmimport_ResourceDrop()
}

//go:wasmimport wasi:http/types@0.2.0-rc-2023-10-18 [resource-drop]future-incoming-response
//go:noescape
func (self FutureIncomingResponse) wasmimport_ResourceDrop()

// Get represents method "get".
//
// option indicates readiness.
// outer result indicates you are allowed to get the
// incoming-response-or-error at most once. subsequent calls after ready
// will return an error here.
// inner result indicates whether the incoming-response was available, or an
// error occured.
//
//	get: func() -> option<result<result<incoming-response, error>>>
//
//go:nosplit
func (self FutureIncomingResponse) Get() cm.Option[cm.OKResult[cm.ErrResult[IncomingResponse, Error], struct{}]] {
	var result cm.Option[cm.OKResult[cm.ErrResult[IncomingResponse, Error], struct{}]]
	self.wasmimport_Get(&result)
	return result
}

//go:wasmimport wasi:http/types@0.2.0-rc-2023-10-18 [method]future-incoming-response.get
//go:noescape
func (self FutureIncomingResponse) wasmimport_Get(result *cm.Option[cm.OKResult[cm.ErrResult[IncomingResponse, Error], struct{}]])

// Subscribe represents method "subscribe".
//
//	subscribe: func() -> pollable
//
//go:nosplit
func (self FutureIncomingResponse) Subscribe() Pollable {
	return self.wasmimport_Subscribe()
}

//go:wasmimport wasi:http/types@0.2.0-rc-2023-10-18 [method]future-incoming-response.subscribe
//go:noescape
func (self FutureIncomingResponse) wasmimport_Subscribe() Pollable

// FutureTrailers represents the resource "wasi:http/types@0.2.0-rc-2023-10-18#future-trailers".
//
//	resource future-trailers
type FutureTrailers cm.Resource

// ResourceDrop represents the Canonical ABI function "resource-drop".
//
// Drops a resource handle.
//
//go:nosplit
func (self FutureTrailers) ResourceDrop() {
	self.wasmimport_ResourceDrop()
}

//go:wasmimport wasi:http/types@0.2.0-rc-2023-10-18 [resource-drop]future-trailers
//go:noescape
func (self FutureTrailers) wasmimport_ResourceDrop()

// Get represents method "get".
//
// Retrieve reference to trailers, if they are ready.
//
//	get: func() -> option<result<trailers, error>>
//
//go:nosplit
func (self FutureTrailers) Get() cm.Option[cm.ErrResult[Trailers, Error]] {
	var result cm.Option[cm.ErrResult[Trailers, Error]]
	self.wasmimport_Get(&result)
	return result
}

//go:wasmimport wasi:http/types@0.2.0-rc-2023-10-18 [method]future-trailers.get
//go:noescape
func (self FutureTrailers) wasmimport_Get(result *cm.Option[cm.ErrResult[Trailers, Error]])

// Subscribe represents method "subscribe".
//
// Pollable that resolves when the body has been fully read, and the trailers
// are ready to be consumed.
//
//	subscribe: func() -> pollable
//
//go:nosplit
func (self FutureTrailers) Subscribe() Pollable {
	return self.wasmimport_Subscribe()
}

//go:wasmimport wasi:http/types@0.2.0-rc-2023-10-18 [method]future-trailers.subscribe
//go:noescape
func (self FutureTrailers) wasmimport_Subscribe() Pollable

// Headers represents the resource "wasi:http/types@0.2.0-rc-2023-10-18#fields".
//
// See [Fields] for more information.
type Headers = Fields

// IncomingBody represents the resource "wasi:http/types@0.2.0-rc-2023-10-18#incoming-body".
//
//	resource incoming-body
type IncomingBody cm.Resource

// ResourceDrop represents the Canonical ABI function "resource-drop".
//
// Drops a resource handle.
//
//go:nosplit
func (self IncomingBody) ResourceDrop() {
	self.wasmimport_ResourceDrop()
}

//go:wasmimport wasi:http/types@0.2.0-rc-2023-10-18 [resource-drop]incoming-body
//go:noescape
func (self IncomingBody) wasmimport_ResourceDrop()

// IncomingBodyFinish represents static function "finish".
//
// takes ownership of incoming-body. this will trap if the
// incoming-body-stream child is still alive!
//
//	finish: static func(this: incoming-body) -> future-trailers
//
//go:nosplit
func IncomingBodyFinish(this IncomingBody) FutureTrailers {
	return wasmimport_IncomingBodyFinish(this)
}

//go:wasmimport wasi:http/types@0.2.0-rc-2023-10-18 [static]incoming-body.finish
//go:noescape
func wasmimport_IncomingBodyFinish(this IncomingBody) FutureTrailers

// Stream represents method "stream".
//
// returned input-stream is a child - the implementation may trap if
// incoming-body is dropped (or consumed by call to
// incoming-body-finish) before the input-stream is dropped.
// May be called at most once. returns error if called additional times.
//
//	stream: func() -> result<input-stream>
//
//go:nosplit
func (self IncomingBody) Stream() cm.OKResult[InputStream, struct{}] {
	var result cm.OKResult[InputStream, struct{}]
	self.wasmimport_Stream(&result)
	return result
}

//go:wasmimport wasi:http/types@0.2.0-rc-2023-10-18 [method]incoming-body.stream
//go:noescape
func (self IncomingBody) wasmimport_Stream(result *cm.OKResult[InputStream, struct{}])

// IncomingRequest represents the resource "wasi:http/types@0.2.0-rc-2023-10-18#incoming-request".
//
// The following block defines the `incoming-request` and `outgoing-request`
// resource types that correspond to HTTP standard Requests.
//
// The `consume` and `write` methods may only be called once (and
// return failure thereafter).
//
//	resource incoming-request
type IncomingRequest cm.Resource

// ResourceDrop represents the Canonical ABI function "resource-drop".
//
// Drops a resource handle.
//
//go:nosplit
func (self IncomingRequest) ResourceDrop() {
	self.wasmimport_ResourceDrop()
}

//go:wasmimport wasi:http/types@0.2.0-rc-2023-10-18 [resource-drop]incoming-request
//go:noescape
func (self IncomingRequest) wasmimport_ResourceDrop()

// Authority represents method "authority".
//
//	authority: func() -> option<string>
//
//go:nosplit
func (self IncomingRequest) Authority() cm.Option[string] {
	var result cm.Option[string]
	self.wasmimport_Authority(&result)
	return result
}

//go:wasmimport wasi:http/types@0.2.0-rc-2023-10-18 [method]incoming-request.authority
//go:noescape
func (self IncomingRequest) wasmimport_Authority(result *cm.Option[string])

// Consume represents method "consume".
//
// Returns the input-stream child at most once.
//
// If called more than once, subsequent calls return an error.
//
//	consume: func() -> result<incoming-body>
//
//go:nosplit
func (self IncomingRequest) Consume() cm.OKResult[IncomingBody, struct{}] {
	var result cm.OKResult[IncomingBody, struct{}]
	self.wasmimport_Consume(&result)
	return result
}

//go:wasmimport wasi:http/types@0.2.0-rc-2023-10-18 [method]incoming-request.consume
//go:noescape
func (self IncomingRequest) wasmimport_Consume(result *cm.OKResult[IncomingBody, struct{}])

// Headers represents method "headers".
//
//	headers: func() -> headers
//
//go:nosplit
func (self IncomingRequest) Headers() Headers {
	return self.wasmimport_Headers()
}

//go:wasmimport wasi:http/types@0.2.0-rc-2023-10-18 [method]incoming-request.headers
//go:noescape
func (self IncomingRequest) wasmimport_Headers() Headers

// Method represents method "method".
//
//	method: func() -> method
//
//go:nosplit
func (self IncomingRequest) Method() Method {
	var result Method
	self.wasmimport_Method(&result)
	return result
}

//go:wasmimport wasi:http/types@0.2.0-rc-2023-10-18 [method]incoming-request.method
//go:noescape
func (self IncomingRequest) wasmimport_Method(result *Method)

// PathWithQuery represents method "path-with-query".
//
//	path-with-query: func() -> option<string>
//
//go:nosplit
func (self IncomingRequest) PathWithQuery() cm.Option[string] {
	var result cm.Option[string]
	self.wasmimport_PathWithQuery(&result)
	return result
}

//go:wasmimport wasi:http/types@0.2.0-rc-2023-10-18 [method]incoming-request.path-with-query
//go:noescape
func (self IncomingRequest) wasmimport_PathWithQuery(result *cm.Option[string])

// Scheme represents method "scheme".
//
//	scheme: func() -> option<scheme>
//
//go:nosplit
func (self IncomingRequest) Scheme() cm.Option[Scheme] {
	var result cm.Option[Scheme]
	self.wasmimport_Scheme(&result)
	return result
}

//go:wasmimport wasi:http/types@0.2.0-rc-2023-10-18 [method]incoming-request.scheme
//go:noescape
func (self IncomingRequest) wasmimport_Scheme(result *cm.Option[Scheme])

// IncomingResponse represents the resource "wasi:http/types@0.2.0-rc-2023-10-18#incoming-response".
//
// The following block defines the `incoming-response` and `outgoing-response`
// resource types that correspond to HTTP standard Responses.
//
// The `consume` and `write` methods may only be called once (and return failure thereafter).
//
//	resource incoming-response
type IncomingResponse cm.Resource

// ResourceDrop represents the Canonical ABI function "resource-drop".
//
// Drops a resource handle.
//
//go:nosplit
func (self IncomingResponse) ResourceDrop() {
	self.wasmimport_ResourceDrop()
}

//go:wasmimport wasi:http/types@0.2.0-rc-2023-10-18 [resource-drop]incoming-response
//go:noescape
func (self IncomingResponse) wasmimport_ResourceDrop()

// Consume represents method "consume".
//
// May be called at most once. returns error if called additional times.
// TODO: make incoming-request-consume work the same way, giving a child
// incoming-body.
//
//	consume: func() -> result<incoming-body>
//
//go:nosplit
func (self IncomingResponse) Consume() cm.OKResult[IncomingBody, struct{}] {
	var result cm.OKResult[IncomingBody, struct{}]
	self.wasmimport_Consume(&result)
	return result
}

//go:wasmimport wasi:http/types@0.2.0-rc-2023-10-18 [method]incoming-response.consume
//go:noescape
func (self IncomingResponse) wasmimport_Consume(result *cm.OKResult[IncomingBody, struct{}])

// Headers represents method "headers".
//
//	headers: func() -> headers
//
//go:nosplit
func (self IncomingResponse) Headers() Headers {
	return self.wasmimport_Headers()
}

//go:wasmimport wasi:http/types@0.2.0-rc-2023-10-18 [method]incoming-response.headers
//go:noescape
func (self IncomingResponse) wasmimport_Headers() Headers

// Status represents method "status".
//
//	status: func() -> status-code
//
//go:nosplit
func (self IncomingResponse) Status() StatusCode {
	return self.wasmimport_Status()
}

//go:wasmimport wasi:http/types@0.2.0-rc-2023-10-18 [method]incoming-response.status
//go:noescape
func (self IncomingResponse) wasmimport_Status() StatusCode

// InputStream represents the resource "wasi:io/streams@0.2.0-rc-2023-10-18#input-stream".
//
// See [streams.InputStream] for more information.
type InputStream = streams.InputStream

// Method represents the variant "wasi:http/types@0.2.0-rc-2023-10-18#method".
//
// This type corresponds to HTTP standard Methods.
//
//	variant method {
//		get,
//		head,
//		post,
//		put,
//		delete,
//		connect,
//		options,
//		trace,
//		patch,
//		other(string),
//	}
type Method cm.Variant[uint8, string, string]

// MethodGet returns a [Method] of case "get".
func MethodGet() Method {
	var data struct{}
	return cm.New[Method](0, data)
}

// Get returns true if [Method] represents the variant case "get".
func (self *Method) Get() bool {
	return cm.Tag(self) == 0
}

// MethodHead returns a [Method] of case "head".
func MethodHead() Method {
	var data struct{}
	return cm.New[Method](1, data)
}

// Head returns true if [Method] represents the variant case "head".
func (self *Method) Head() bool {
	return cm.Tag(self) == 1
}

// MethodPost returns a [Method] of case "post".
func MethodPost() Method {
	var data struct{}
	return cm.New[Method](2, data)
}

// Post returns true if [Method] represents the variant case "post".
func (self *Method) Post() bool {
	return cm.Tag(self) == 2
}

// MethodPut returns a [Method] of case "put".
func MethodPut() Method {
	var data struct{}
	return cm.New[Method](3, data)
}

// Put returns true if [Method] represents the variant case "put".
func (self *Method) Put() bool {
	return cm.Tag(self) == 3
}

// MethodDelete returns a [Method] of case "delete".
func MethodDelete() Method {
	var data struct{}
	return cm.New[Method](4, data)
}

// Delete returns true if [Method] represents the variant case "delete".
func (self *Method) Delete() bool {
	return cm.Tag(self) == 4
}

// MethodConnect returns a [Method] of case "connect".
func MethodConnect() Method {
	var data struct{}
	return cm.New[Method](5, data)
}

// Connect returns true if [Method] represents the variant case "connect".
func (self *Method) Connect() bool {
	return cm.Tag(self) == 5
}

// MethodOptions returns a [Method] of case "options".
func MethodOptions() Method {
	var data struct{}
	return cm.New[Method](6, data)
}

// Options returns true if [Method] represents the variant case "options".
func (self *Method) Options() bool {
	return cm.Tag(self) == 6
}

// MethodTrace returns a [Method] of case "trace".
func MethodTrace() Method {
	var data struct{}
	return cm.New[Method](7, data)
}

// Trace returns true if [Method] represents the variant case "trace".
func (self *Method) Trace() bool {
	return cm.Tag(self) == 7
}

// MethodPatch returns a [Method] of case "patch".
func MethodPatch() Method {
	var data struct{}
	return cm.New[Method](8, data)
}

// Patch returns true if [Method] represents the variant case "patch".
func (self *Method) Patch() bool {
	return cm.Tag(self) == 8
}

// MethodOther returns a [Method] of case "other".
func MethodOther(data string) Method {
	return cm.New[Method](9, data)
}

// Other returns a non-nil *[string] if [Method] represents the variant case "other".
func (self *Method) Other() *string {
	return cm.Case[string](self, 9)
}

// OutgoingBody represents the resource "wasi:http/types@0.2.0-rc-2023-10-18#outgoing-body".
//
//	resource outgoing-body
type OutgoingBody cm.Resource

// ResourceDrop represents the Canonical ABI function "resource-drop".
//
// Drops a resource handle.
//
//go:nosplit
func (self OutgoingBody) ResourceDrop() {
	self.wasmimport_ResourceDrop()
}

//go:wasmimport wasi:http/types@0.2.0-rc-2023-10-18 [resource-drop]outgoing-body
//go:noescape
func (self OutgoingBody) wasmimport_ResourceDrop()

// OutgoingBodyFinish represents static function "finish".
//
// Finalize an outgoing body, optionally providing trailers. This must be
// called to signal that the response is complete. If the `outgoing-body` is
// dropped without calling `outgoing-body-finalize`, the implementation
// should treat the body as corrupted.
//
//	finish: static func(this: outgoing-body, trailers: option<trailers>)
//
//go:nosplit
func OutgoingBodyFinish(this OutgoingBody, trailers cm.Option[Trailers]) {
	wasmimport_OutgoingBodyFinish(this, trailers)
}

//go:wasmimport wasi:http/types@0.2.0-rc-2023-10-18 [static]outgoing-body.finish
//go:noescape
func wasmimport_OutgoingBodyFinish(this OutgoingBody, trailers cm.Option[Trailers])

// Write represents method "write".
//
// Will give the child output-stream at most once. subsequent calls will
// return an error.
//
//	write: func() -> result<output-stream>
//
//go:nosplit
func (self OutgoingBody) Write() cm.OKResult[OutputStream, struct{}] {
	var result cm.OKResult[OutputStream, struct{}]
	self.wasmimport_Write(&result)
	return result
}

//go:wasmimport wasi:http/types@0.2.0-rc-2023-10-18 [method]outgoing-body.write
//go:noescape
func (self OutgoingBody) wasmimport_Write(result *cm.OKResult[OutputStream, struct{}])

// OutgoingRequest represents the resource "wasi:http/types@0.2.0-rc-2023-10-18#outgoing-request".
//
//	resource outgoing-request
type OutgoingRequest cm.Resource

// ResourceDrop represents the Canonical ABI function "resource-drop".
//
// Drops a resource handle.
//
//go:nosplit
func (self OutgoingRequest) ResourceDrop() {
	self.wasmimport_ResourceDrop()
}

//go:wasmimport wasi:http/types@0.2.0-rc-2023-10-18 [resource-drop]outgoing-request
//go:noescape
func (self OutgoingRequest) wasmimport_ResourceDrop()

// NewOutgoingRequest represents constructor "[constructor]outgoing-request".
//
//	[constructor]outgoing-request(method: method, path-with-query: option<string>,
//	scheme: option<scheme>, authority: option<string>, headers: borrow<headers>)
//
//go:nosplit
func NewOutgoingRequest(method Method, pathWithQuery cm.Option[string], scheme cm.Option[Scheme], authority cm.Option[string], headers Headers) OutgoingRequest {
	return wasmimport_NewOutgoingRequest(method, pathWithQuery, scheme, authority, headers)
}

//go:wasmimport wasi:http/types@0.2.0-rc-2023-10-18 [constructor]outgoing-request
//go:noescape
func wasmimport_NewOutgoingRequest(method Method, pathWithQuery cm.Option[string], scheme cm.Option[Scheme], authority cm.Option[string], headers Headers) OutgoingRequest

// Write represents method "write".
//
// Will return the outgoing-body child at most once.
//
// If called more than once, subsequent calls return an error.
//
//	write: func() -> result<outgoing-body>
//
//go:nosplit
func (self OutgoingRequest) Write() cm.OKResult[OutgoingBody, struct{}] {
	var result cm.OKResult[OutgoingBody, struct{}]
	self.wasmimport_Write(&result)
	return result
}

//go:wasmimport wasi:http/types@0.2.0-rc-2023-10-18 [method]outgoing-request.write
//go:noescape
func (self OutgoingRequest) wasmimport_Write(result *cm.OKResult[OutgoingBody, struct{}])

// OutgoingResponse represents the resource "wasi:http/types@0.2.0-rc-2023-10-18#outgoing-response".
//
//	resource outgoing-response
type OutgoingResponse cm.Resource

// ResourceDrop represents the Canonical ABI function "resource-drop".
//
// Drops a resource handle.
//
//go:nosplit
func (self OutgoingResponse) ResourceDrop() {
	self.wasmimport_ResourceDrop()
}

//go:wasmimport wasi:http/types@0.2.0-rc-2023-10-18 [resource-drop]outgoing-response
//go:noescape
func (self OutgoingResponse) wasmimport_ResourceDrop()

// NewOutgoingResponse represents constructor "[constructor]outgoing-response".
//
//	[constructor]outgoing-response(status-code: status-code, headers: borrow<headers>)
//
//go:nosplit
func NewOutgoingResponse(statusCode StatusCode, headers Headers) OutgoingResponse {
	return wasmimport_NewOutgoingResponse(statusCode, headers)
}

//go:wasmimport wasi:http/types@0.2.0-rc-2023-10-18 [constructor]outgoing-response
//go:noescape
func wasmimport_NewOutgoingResponse(statusCode StatusCode, headers Headers) OutgoingResponse

// Write represents method "write".
//
// Will give the child outgoing-response at most once. subsequent calls will
// return an error.
//
//	write: func() -> result<outgoing-body>
//
//go:nosplit
func (self OutgoingResponse) Write() cm.OKResult[OutgoingBody, struct{}] {
	var result cm.OKResult[OutgoingBody, struct{}]
	self.wasmimport_Write(&result)
	return result
}

//go:wasmimport wasi:http/types@0.2.0-rc-2023-10-18 [method]outgoing-response.write
//go:noescape
func (self OutgoingResponse) wasmimport_Write(result *cm.OKResult[OutgoingBody, struct{}])

// OutputStream represents the resource "wasi:io/streams@0.2.0-rc-2023-10-18#output-stream".
//
// See [streams.OutputStream] for more information.
type OutputStream = streams.OutputStream

// Pollable represents the resource "wasi:io/poll@0.2.0-rc-2023-10-18#pollable".
//
// See [poll.Pollable] for more information.
type Pollable = poll.Pollable

// RequestOptions represents the record "wasi:http/types@0.2.0-rc-2023-10-18#request-options".
//
// Additional optional parameters that can be set when making a request.
//
//	record request-options {
//		connect-timeout-ms: option<u32>,
//		first-byte-timeout-ms: option<u32>,
//		between-bytes-timeout-ms: option<u32>,
//	}
type RequestOptions struct {
	// The following timeouts are specific to the HTTP protocol and work
	// independently of the overall timeouts passed to `io.poll.poll-list`.
	// The timeout for the initial connect.
	ConnectTimeoutMs cm.Option[uint32]

	// The timeout for receiving the first byte of the response body.
	FirstByteTimeoutMs cm.Option[uint32]

	// The timeout for receiving the next chunk of bytes in the response body
	// stream.
	BetweenBytesTimeoutMs cm.Option[uint32]
}

// ResponseOutparam represents the resource "wasi:http/types@0.2.0-rc-2023-10-18#response-outparam".
//
// The following block defines a special resource type used by the
// `wasi:http/incoming-handler` interface. When resource types are added, this
// block can be replaced by a proper `resource response-outparam { ... }`
// definition. Later, with Preview3, the need for an outparam goes away entirely
// (the `wasi:http/handler` interface used for both incoming and outgoing can
// simply return a `stream`).
//
//	resource response-outparam
type ResponseOutparam cm.Resource

// ResourceDrop represents the Canonical ABI function "resource-drop".
//
// Drops a resource handle.
//
//go:nosplit
func (self ResponseOutparam) ResourceDrop() {
	self.wasmimport_ResourceDrop()
}

//go:wasmimport wasi:http/types@0.2.0-rc-2023-10-18 [resource-drop]response-outparam
//go:noescape
func (self ResponseOutparam) wasmimport_ResourceDrop()

// ResponseOutparamSet represents static function "set".
//
//	set: static func(param: response-outparam, response: result<outgoing-response,
//	error>)
//
//go:nosplit
func ResponseOutparamSet(param ResponseOutparam, response cm.ErrResult[OutgoingResponse, Error]) {
	wasmimport_ResponseOutparamSet(param, response)
}

//go:wasmimport wasi:http/types@0.2.0-rc-2023-10-18 [static]response-outparam.set
//go:noescape
func wasmimport_ResponseOutparamSet(param ResponseOutparam, response cm.ErrResult[OutgoingResponse, Error])

// Scheme represents the variant "wasi:http/types@0.2.0-rc-2023-10-18#scheme".
//
// This type corresponds to HTTP standard Related Schemes.
//
//	variant scheme {
//		HTTP,
//		HTTPS,
//		other(string),
//	}
type Scheme cm.Variant[uint8, string, string]

// SchemeHTTP returns a [Scheme] of case "HTTP".
func SchemeHTTP() Scheme {
	var data struct{}
	return cm.New[Scheme](0, data)
}

// HTTP returns true if [Scheme] represents the variant case "HTTP".
func (self *Scheme) HTTP() bool {
	return cm.Tag(self) == 0
}

// SchemeHTTPS returns a [Scheme] of case "HTTPS".
func SchemeHTTPS() Scheme {
	var data struct{}
	return cm.New[Scheme](1, data)
}

// HTTPS returns true if [Scheme] represents the variant case "HTTPS".
func (self *Scheme) HTTPS() bool {
	return cm.Tag(self) == 1
}

// SchemeOther returns a [Scheme] of case "other".
func SchemeOther(data string) Scheme {
	return cm.New[Scheme](2, data)
}

// Other returns a non-nil *[string] if [Scheme] represents the variant case "other".
func (self *Scheme) Other() *string {
	return cm.Case[string](self, 2)
}

// StatusCode represents the type "wasi:http/types@0.2.0-rc-2023-10-18#status-code".
//
// This type corresponds to the HTTP standard Status Code.
//
//	type status-code = u16
type StatusCode uint16

// Trailers represents the resource "wasi:http/types@0.2.0-rc-2023-10-18#fields".
//
// See [Fields] for more information.
type Trailers = Fields

type Interface interface {
	NewFields(entries cm.List[cm.Tuple[string, cm.List[uint8]]]) Fields
	NewOutgoingRequest(method Method, pathWithQuery cm.Option[string], scheme cm.Option[Scheme], authority cm.Option[string], headers Headers) OutgoingRequest
	NewOutgoingResponse(statusCode StatusCode, headers Headers) OutgoingResponse
	IncomingBodyFinish(this IncomingBody) FutureTrailers
	OutgoingBodyFinish(this OutgoingBody, trailers cm.Option[Trailers])
	ResponseOutparamSet(param ResponseOutparam, response cm.ErrResult[OutgoingResponse, Error])
}

var instance Interface
