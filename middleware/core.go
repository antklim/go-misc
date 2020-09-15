package middleware

import (
	"fmt"
	"strings"
)

////////////////////////////////////////////////////////////////////////////////
// Core functions package
////////////////////////////////////////////////////////////////////////////////

func echo(s string) (string, error) {
	return s, nil
}

func lowercase(s string) (string, error) {
	return strings.ToLower(s), nil
}

func uppercase(s string) (string, error) {
	return strings.ToUpper(s), nil
}

func swapCase(s string) (string, error) {
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

func reverse(s string) (string, error) {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes), nil
}

func strWrapper(lw, rw string) func(s string) (string, error) {
	return func(s string) (string, error) {
		var b strings.Builder
		if _, err := b.WriteString(lw); err != nil {
			return "", err
		}

		if _, err := b.WriteString(s); err != nil {
			return "", err
		}

		if _, err := b.WriteString(rw); err != nil {
			return "", err
		}

		return b.String(), nil
	}
}

func defStrWrapper(s string) (string, error) {
	return strWrapper("[", "]")(s)
}

func validate(s string) (string, error) {
	n := len(s)
	if n < 1 || n > 100 {
		return "", fmt.Errorf("invalid string length %d", n)
	}

	return s, nil
}
