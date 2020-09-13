package middleware

import (
	"fmt"
	"strings"
)

/**
TODO: Wrapper
			Wrapper(a, b, c)
				Execution flow:
					c (before invoking handler) ->
					b  (before invoking handler) ->
					a ->
					b deferred functions ->
					c deferred functions

TODO: Chain
			Chain(a, b, c)
				Execution flow:
					a -> b -> c
*/

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

type wrap func(Handler) Handler

// Wrap chains middleware calls.
func Wrap(handlers ...Handler) Handler {
	var w wrap
	var r Handler
	for _, h := range handlers {
		r = w(h)
	}
	return r
}

// Echo echo middleware.
type Echo struct{}

// ServeString implements Handler interface.
func (e Echo) ServeString(s string) (string, error) {
	return s, nil
}

// Validate is a string validation middleware.
func Validate(s string) (string, error) {
	n := len(s)
	if n < 1 || n > 100 {
		return "", fmt.Errorf("invalid string length %d", n)
	}

	return s, nil
}

// Lowercase is a string lowercase middleware.
func Lowercase() HandlerFunc {
	return HandlerFunc(func(s string) (string, error) {
		return strings.ToLower(s), nil
	})
}

// Uppercase is a string uppercase middleware.
func Uppercase() HandlerFunc {
	return HandlerFunc(func(s string) (string, error) {
		return strings.ToUpper(s), nil
	})
}

// SwapCase is a string to swap character case middleware.
func SwapCase(s string) (string, error) {
	sc := func(r rune) rune {
		switch {
		case r >= 'a' && r <= 'z', r >= 'A' && r <= 'Z':
			return r ^ 32
		default:
			return r
		}
	}
	return strings.Map(sc, s), nil
}

// Reverse is a string reverse middleware.
func Reverse(s string) (string, error) {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes), nil
}

// DefaultStrWrapper is a default string wrapper middleware.
func DefaultStrWrapper(h Handler) Handler {
	return HandlerFunc(func(s string) (string, error) {
		var b strings.Builder
		if _, err := b.WriteRune('['); err != nil {
			return "", err
		}

		ss, err := h.ServeString(s)
		if err != nil {
			return "", err
		}

		if _, err := b.WriteString(ss); err != nil {
			return "", err
		}

		if _, err := b.WriteRune(']'); err != nil {
			return "", err
		}

		return b.String(), nil
	})
}

// StrWrapper wraps string with left and right strings.
func StrWrapper(h Handler, lw, rw string) Handler {
	return HandlerFunc(func(s string) (string, error) {
		var b strings.Builder
		if _, err := b.WriteString(lw); err != nil {
			return "", err
		}

		ss, err := h.ServeString(s)
		if err != nil {
			return "", err
		}

		if _, err := b.WriteString(ss); err != nil {
			return "", err
		}

		if _, err := b.WriteString(rw); err != nil {
			return "", err
		}

		return b.String(), nil
	})
}

// Handle entry point to middleware showcase.
func Handle(s string, handler Handler) (string, error) {
	return handler.ServeString(s)
}
