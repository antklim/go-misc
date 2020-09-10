package create_test

import (
	"testing"

	"github.com/antklim/go-misc/create"
	"github.com/stretchr/testify/assert"
)

func TestTextMessageBuilder(t *testing.T) {
	mb := create.NewTextMessageBuilder()
	mb.AddRecipient("A", "B").
		AddSubject("Welcome to builder").
		AddGreeting("Dear A and B").
		AddBody("Nice text").
		AddFooter("With nice footer!")

	assert.Equal(t, textMessageExpected, mb.GetMessage())
}

func TestJsonMessageBuilder(t *testing.T) {
	mb := create.NewJsonMessageBuilder()
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
	mb := create.GetBuilder(create.Json)
	mb.AddBody("Hi there!")

	msg, err := mb.(*create.JsonMessage).GetMessage()
	assert.NoError(t, err)
	assert.Equal(t, `{"body":"Hi there!"}`, string(msg))
}
