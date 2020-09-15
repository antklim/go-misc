package middleware

/**
Wrap(a, b, c)
	Execution flow:
		c (before invoking handler) ->
		b  (before invoking handler) ->
		a ->
		b deferred functions ->
		c deferred functions

Chain(a, b, c)
	Execution flow:
		a -> b -> c
*/

// Chain chains handlers calls and passes output of handler to input of the next handler.
func Chain(hh ...Handler) Handler {
	return HandlerFunc(func(s string) (string, error) {
		var err error
		for _, h := range hh {
			s, err = h.Serve(s)
			if err != nil {
				return "", err
			}
		}
		return s, nil
	})
}

// Wrap wraps handlers in a way the last HadnlerWrap executed first.
func Wrap(h Handler, hww ...HandlerWrap) Handler {
	for _, hw := range hww {
		h = hw(h)
	}
	return h
}

// Handle entry point to middleware showcase.
func Handle(s string, handler Handler) (string, error) {
	return handler.Serve(s)
}
