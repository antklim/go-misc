package middleware_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/antklim/go-misc/middleware"
)

func TestEcho(t *testing.T) {
	mw := middleware.Echo{}
	actual, err := middleware.Handle("hi echo", mw)
	assert.NoError(t, err)
	assert.Equal(t, "hi echo", actual)
}

func TestValidate(t *testing.T) {
	mw := middleware.HandlerFunc(middleware.Validate)
	for _, tC := range validateTests {
		t.Run(tC.desc, func(t *testing.T) {
			_, err := middleware.Handle(tC.s, mw)
			assert.EqualError(t, err, tC.expected)
		})
	}
}

func TestSwapCase(t *testing.T) {
	mw := middleware.HandlerFunc(middleware.SwapCase)
	actual, err := middleware.Handle("aBcD1234", mw)
	assert.NoError(t, err)
	assert.Equal(t, "AbCd1234", actual)
}

func TestReverse(t *testing.T) {
	mw := middleware.HandlerFunc(middleware.Reverse)
	actual, err := middleware.Handle("abcd1234", mw)
	assert.NoError(t, err)
	assert.Equal(t, "4321dcba", actual)
}

func TestWrap(t *testing.T) {
	mw := middleware.HandlerFunc(middleware.Wrap)
	actual, err := middleware.Handle("hi wrap", mw)
	assert.NoError(t, err)
	assert.Equal(t, "[hi wrap]", actual)
}
