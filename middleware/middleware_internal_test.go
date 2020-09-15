package middleware

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCoreEcho(t *testing.T) {
	actual, err := echo("hello")
	require.NoError(t, err)
	assert.Equal(t, "hello", actual)
}

func TestCoreLowercase(t *testing.T) {
	actual, err := lowercase("HEllo")
	require.NoError(t, err)
	assert.Equal(t, "hello", actual)
}

func TestCoreUppercase(t *testing.T) {
	actual, err := uppercase("HEllo")
	require.NoError(t, err)
	assert.Equal(t, "HELLO", actual)
}

func TestCoreSwapCase(t *testing.T) {
	actual, err := swapCase("HEllo")
	require.NoError(t, err)
	assert.Equal(t, "heLLO", actual)
}

func TestCoreReverse(t *testing.T) {
	actual, err := reverse("HEllo")
	require.NoError(t, err)
	assert.Equal(t, "ollEH", actual)
}

func TestCoreDefStrWrapper(t *testing.T) {
	actual, err := defStrWrapper("HEllo")
	require.NoError(t, err)
	assert.Equal(t, "[HEllo]", actual)
}

func TestCoreStrWrapper(t *testing.T) {
	actual, err := strWrapper(">", "<")("HEllo")
	require.NoError(t, err)
	assert.Equal(t, ">HEllo<", actual)
}
