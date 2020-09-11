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

func TestWrap(t *testing.T) {
	mw := middleware.HandlerFunc(middleware.Wrap)
	actual, err := middleware.Handle("hi wrap", mw)
	assert.NoError(t, err)
	assert.Equal(t, "[hi wrap]", actual)
}
