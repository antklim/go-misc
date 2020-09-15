package middleware

import (
	"fmt"
)

// Handler interface that should be implemented by middlewares.
type Handler interface {
	Serve(s string) (string, error)
}

// HandlerFunc adapter to allow use of ordinary functions as middleware.
type HandlerFunc func(s string) (string, error)

// HandlerWrap adapter to wrap handler into other handler.
type HandlerWrap func(Handler) Handler

// Serve implements Handler interface.
func (f HandlerFunc) Serve(s string) (string, error) {
	return f(s)
}

// EchoH is echo handler.
func EchoH() HandlerFunc {
	return HandlerFunc(echo)
}

// LowercaseH is lower case handler.
func LowercaseH() HandlerFunc {
	return HandlerFunc(lowercase)
}

// UppercaseH is upper case handler.
func UppercaseH() HandlerFunc {
	return HandlerFunc(uppercase)
}

// SwapCaseH is swap case handler.
func SwapCaseH() HandlerFunc {
	return HandlerFunc(swapCase)
}

// ReverseH is reverse string handler.
func ReverseH() HandlerFunc {
	return HandlerFunc(reverse)
}

// StrWrapperH is string wrapper handler.
func StrWrapperH(lw, rw string) HandlerFunc {
	return HandlerFunc(strWrapper(lw, rw))
}

// StrWrapperW is string wrapper handler.
func StrWrapperW(lw, rw string) HandlerWrap {
	return HandlerWrap(func(h Handler) Handler {

		return HandlerFunc(func(s string) (string, error) {
			fmt.Println("StrWrapperW before handler")
			s = fmt.Sprintf("%s%s", lw, s)

			defer func() {
				fmt.Println("StrWrapperW after handler")
				// The following line does not interfere/change s in outer scope
				s = fmt.Sprintf("%s%s", s, rw)
			}()

			s, err := h.Serve(s)
			s = fmt.Sprintf("%s%s", s, rw)
			return s, err
		})

	})
}

// DefStrWrapperH is default string wrapper handler.
func DefStrWrapperH() HandlerFunc {
	return HandlerFunc(defStrWrapper)
}

// ValidateH is validate handler.
func ValidateH() HandlerFunc {
	return HandlerFunc(validate)
}

// ValidateW validate handler wrapper.
func ValidateW() HandlerWrap {
	return HandlerWrap(func(h Handler) Handler {

		return HandlerFunc(func(s string) (string, error) {
			fmt.Println("ValidateW before handler")
			defer fmt.Println("ValidateW after handler")
			return h.Serve(s)
		})

	})
}
