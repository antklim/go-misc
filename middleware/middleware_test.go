package middleware_test

import (
	"testing"

	mw "github.com/antklim/go-misc/middleware"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChain(t *testing.T) {
	chain := mw.Chain(mw.SwapCaseH(), mw.StrWrapperH("A", "b"), mw.ReverseH(), mw.EchoH())
	actual, err := mw.Handle("HeLLo", chain)
	require.NoError(t, err)
	assert.Equal(t, "bOllEhA", actual)
}

func TestWrap(t *testing.T) {
	wrap := mw.Wrap(mw.SwapCaseH(), mw.StrWrapperW("CC", "n"), mw.ValidateW())
	actual, err := mw.Handle("HeLLo", wrap)
	require.NoError(t, err)
	assert.Equal(t, "cchEllOn", actual)
}
