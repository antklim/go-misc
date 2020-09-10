package create_test

import (
	"testing"

	"github.com/antklim/go-misc/create"
	"github.com/stretchr/testify/assert"
)

func TestFacebookConnector(t *testing.T) {
	c := create.NewFacebookConnector("user", "password")
	res := c.Login()
	assert.Equal(t, "logged in to facebook", res)

	res = c.CreatePost("hey")
	assert.Equal(t, "posting [hey] to facebook", res)

	res = c.LogOut()
	assert.Equal(t, "logged out from facebook", res)
}

func TestTwitterConnector(t *testing.T) {
	c := create.NewTwitterConnector("apikey")
	res := c.Login()
	assert.Equal(t, "logged in to twitter", res)

	res = c.CreatePost("hey111")
	assert.Equal(t, "posting [hey111] to twitter", res)

	res = c.LogOut()
	assert.Equal(t, "logged out from twitter", res)
}

func TestPoster(t *testing.T) {
	fb := create.NewPoster(create.NewFacebookPoster("user", "login"))
	tw := create.NewPoster(create.NewTwitterPoster("abc123"))
	fp := fb.Post("blah")
	tp := tw.Post("booo")

	assert.Equal(t, "posting [blah] to facebook", fp)
	assert.Equal(t, "posting [booo] to twitter", tp)
}
