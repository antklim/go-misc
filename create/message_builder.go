package create

import (
	"encoding/json"
	"fmt"
)

type MessageType int

const (
	Text MessageType = iota + 1
	Json
)

// MessageBuilder ...
type MessageBuilder interface {
	AddRecipient(rs ...string) MessageBuilder
	AddSubject(s string) MessageBuilder
	AddGreeting(g string) MessageBuilder
	AddBody(b string) MessageBuilder
	AddFooter(f string) MessageBuilder
}

func GetBuilder(mt MessageType) MessageBuilder {
	switch mt {
	case Text:
		return NewTextMessageBuilder()
	case Json:
		return NewJsonMessageBuilder()
	default:
		return nil
	}
}

// Message ...
type Message struct {
	Recipients []string `json:"recipients,omitempty"`
	Subject    string   `json:"subject,omitempty"`
	Greeting   string   `json:"greeting,omitempty"`
	Body       string   `json:"body,omitempty"`
	Footer     string   `json:"footer,omitempty"`
}

type TextMessage struct {
	message Message
}

func NewTextMessageBuilder() *TextMessage {
	return &TextMessage{}
}

func (m *TextMessage) AddRecipient(rs ...string) MessageBuilder {
	m.message.Recipients = append(m.message.Recipients, rs...)
	return m
}

func (m *TextMessage) AddSubject(s string) MessageBuilder {
	m.message.Subject = s
	return m
}

func (m *TextMessage) AddGreeting(g string) MessageBuilder {
	m.message.Greeting = g
	return m
}

func (m *TextMessage) AddBody(b string) MessageBuilder {
	m.message.Body = b
	return m
}

func (m *TextMessage) AddFooter(f string) MessageBuilder {
	m.message.Footer = f
	return m
}

func (m TextMessage) GetMessage() string {
	return fmt.Sprintf("recipients: %s\nsubject: %s\n\n\t%s\n%s\n%s",
		m.message.Recipients, m.message.Subject, m.message.Greeting, m.message.Body, m.message.Footer)
}

type JsonMessage struct {
	message Message
}

func NewJsonMessageBuilder() *JsonMessage {
	return &JsonMessage{}
}

func (m *JsonMessage) AddRecipient(rs ...string) MessageBuilder {
	m.message.Recipients = append(m.message.Recipients, rs...)
	return m
}

func (m *JsonMessage) AddSubject(s string) MessageBuilder {
	m.message.Subject = s
	return m
}

func (m *JsonMessage) AddGreeting(g string) MessageBuilder {
	m.message.Greeting = g
	return m
}

func (m *JsonMessage) AddBody(b string) MessageBuilder {
	m.message.Body = b
	return m
}

func (m *JsonMessage) AddFooter(f string) MessageBuilder {
	m.message.Footer = f
	return m
}

func (m JsonMessage) GetMessage() ([]byte, error) {
	return json.Marshal(m.message)
}
