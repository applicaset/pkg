// Package httpx is the contract between services: a JSON request, a JSON response, and one error
// shape for failures caused by the request, not the system. It imports only the standard library
// and pkg/reqid, so no service ends up depending on another.
package httpx

import "net/http"

// Code lists the only ways a request, not the system, can be at fault. Any other failure must never
// be read as an answer.
type Code string

const (
	CodeNotFound     Code = "not_found"
	CodeInvalidInput Code = "invalid_input"
	CodeConflict     Code = "conflict"
	CodeForbidden    Code = "forbidden"
)

// Known reports whether this is a code a client may act on. A code it does not recognise is a
// failure of the system, not a domain answer.
func (c Code) Known() bool {
	switch c {
	case CodeNotFound, CodeInvalidInput, CodeConflict, CodeForbidden:
		return true
	default:
		return false
	}
}

// Status is the HTTP status a code is carried on.
func (c Code) Status() int {
	switch c {
	case CodeNotFound:
		return http.StatusNotFound
	case CodeInvalidInput:
		return http.StatusBadRequest
	case CodeConflict:
		return http.StatusConflict
	case CodeForbidden:
		return http.StatusForbidden
	default:
		return http.StatusInternalServerError
	}
}

// Envelope is the body of every non-2xx response from a service API, and the only shape a client
// accepts as a domain error. Message is written for a visitor.
type Envelope struct {
	Code    Code   `json:"code"`
	Message string `json:"message"`
}

// Error is what a client returns for a well-formed domain failure. Anything else comes back as an
// ordinary wrapped error, so "the service is down" is never read as an answer.
type Error struct {
	Code    Code
	Message string
}

func (e *Error) Error() string {
	return string(e.Code) + ": " + e.Message
}
