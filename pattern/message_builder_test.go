package pattern_test

import (
	"testing"

	"github.com/antklim/go-misc/pattern"
	"github.com/stretchr/testify/assert"
)

func TestTextMessageBuilder(t *testing.T) {
	mb := pattern.NewTextMessageBuilder()
	mb.AddRecipient("A", "B").
		AddSubject("Welcome to builder").
		AddGreeting("Dear A and B").
		AddBody("Nice text").
		AddFooter("With nice footer!")

	assert.Equal(t, textMessageExpected, mb.GetMessage())
}

func TestJsonMessageBuilder(t *testing.T) {
	mb := pattern.NewJsonMessageBuilder()
	mb.AddRecipient("A", "B").
		AddSubject("Welcome to builder").
		AddGreeting("Dear A and B").
		AddBody("Nice text").
		AddFooter("With nice footer!")

	msg, err := mb.GetMessage()
	assert.NoError(t, err)
	assert.Equal(t, jsonMessageExpected, string(msg))
}

func TestGetBuilder(t *testing.T) {
	mb := pattern.GetBuilder(pattern.Json)
	mb.AddBody("Hi there!")

	msg, err := mb.(*pattern.JsonMessage).GetMessage()
	assert.NoError(t, err)
	assert.Equal(t, `{"body":"Hi there!"}`, string(msg))
}
