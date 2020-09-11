package middleware

import (
	"fmt"
)

// Implement middleware that validates string length > 0 & < 100
// Implement middleware that trims string
// Implement middleware that UPPERCASE string
// Implement middleware that lowercase string
// Implement middleware that reverse string
// Chain middlewares

// Handler interface that should be implemented by middlewares.
type Handler interface {
	ServeString(s string) (string, error)
}

// HandlerFunc adapter to allow use of ordinary functions as middleware.
type HandlerFunc func(s string) (string, error)

// ServeString implements Handler interface.
func (f HandlerFunc) ServeString(s string) (string, error) {
	return f(s)
}

// Echo echo middleware.
type Echo struct{}

// ServeString implements Handler interface.
func (e Echo) ServeString(s string) (string, error) {
	return s, nil
}

// Wrap is a wrap middleware.
func Wrap(s string) (string, error) {
	return fmt.Sprintf("[%s]", s), nil
}

// Handle entry point to middleware shocase.
func Handle(s string, handler Handler) (string, error) {
	return handler.ServeString(s)
}
