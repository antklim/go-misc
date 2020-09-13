package middleware_test

import (
	"testing"

	"github.com/antklim/go-misc/middleware"
	"github.com/stretchr/testify/assert"
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

func TestLowerCase(t *testing.T) {
	actual, err := middleware.Handle("aBcD1234", middleware.Lowercase())
	assert.NoError(t, err)
	assert.Equal(t, "abcd1234", actual)
}

func TestUpperCase(t *testing.T) {
	actual, err := middleware.Handle("aBcD1234", middleware.Uppercase())
	assert.NoError(t, err)
	assert.Equal(t, "ABCD1234", actual)
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

func TestDefaultStrWrapper(t *testing.T) {
	echoMw := middleware.Echo{}
	actual, err := middleware.Handle("hi wrap", middleware.DefaultStrWrapper(echoMw))
	assert.NoError(t, err)
	assert.Equal(t, "[hi wrap]", actual)
}

// func TestChain(t *testing.T) {
// 	echoMw := middleware.Echo{}
// 	actual, err := middleware.Handle("Hi Middleware", middleware.Chain(echoMw, ))
// 	assert.NoError(t, err)
// 	assert.Equal(t, "Hi Middleware", actual)
// }
