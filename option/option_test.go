package option_test

import (
	"testing"
	"time"

	"github.com/antklim/go-misc/option"
	"github.com/stretchr/testify/assert"
)

func TestDefaultOptions(t *testing.T) {
	expected := `address: localhost:8080
with cache: false
client timeout: 300ms
max connections: 100
logger provided: false`

	opt := []option.Option{}
	s := option.NewServer(opt...)
	assert.Equal(t, expected, s.Info())
}

func TestCustomOptions(t *testing.T) {
	expected := `address: localhost:8080
with cache: true
client timeout: 100ms
max connections: 500
logger provided: false`

	opt := []option.Option{
		option.MaxConcurrentConnections(500),
		option.ClientTimeout(100 * time.Millisecond),
		option.WithCache(true),
	}
	s := option.NewServer(opt...)
	assert.Equal(t, expected, s.Info())
}
