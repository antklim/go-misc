package create

import (
	"fmt"
)

////////////////////////////////////////////////////////////////////////////////
// Creators
////////////////////////////////////////////////////////////////////////////////

type Poster struct {
	smp SocialMediaPoster
}

func NewPoster(smp SocialMediaPoster) *Poster {
	return &Poster{smp}
}

func (p Poster) Post(post string) string {
	conn := p.smp.GetConnector()

	conn.Login()
	res := conn.CreatePost(post)
	conn.LogOut()
	return res
}

type SocialMediaPoster interface {
	GetConnector() SocialNetworkConnector
}

// FacebookPoster concrete creator supports Facebook.
type FacebookPoster struct {
	login    string
	password string
}

// NewFacebookPoster creates new FacebookPoster.
func NewFacebookPoster(login, password string) *FacebookPoster {
	return &FacebookPoster{login, password}
}

func (p FacebookPoster) GetConnector() SocialNetworkConnector {
	return NewFacebookConnector(p.login, p.password)
}

// TwitterPoster concrete creator supports Twitter.
type TwitterPoster struct {
	apikey string
}

// NewTwitterPoster creates new TwitterPoster.
func NewTwitterPoster(apikey string) *TwitterPoster {
	return &TwitterPoster{apikey}
}

func (p TwitterPoster) GetConnector() SocialNetworkConnector {
	return NewTwitterConnector(p.apikey)
}

////////////////////////////////////////////////////////////////////////////////
// Products
////////////////////////////////////////////////////////////////////////////////

type SocialNetworkConnector interface {
	Login() string
	LogOut() string
	CreatePost(post string) string
}

type FacebookConnector struct {
	login    string
	password string
	// conn facebook api connection
}

func NewFacebookConnector(login, password string) SocialNetworkConnector {
	return &FacebookConnector{login, password}
}

func (c *FacebookConnector) Login() string {
	// establish connection with facebook api
	// conn, err := facebookApi.Connect(c.login, c.password)
	// c.conn = conn
	return "logged in to facebook"
}

func (c FacebookConnector) LogOut() string {
	// c.conn.Logout()
	return "logged out from facebook"
}

func (c FacebookConnector) CreatePost(post string) string {
	// c.conn.CreatePost(post)
	return fmt.Sprintf("posting [%s] to facebook", post)
}

type TwitterConnector struct {
	apikey string
	// conn twitter api connection
}

func NewTwitterConnector(apikey string) SocialNetworkConnector {
	return &TwitterConnector{apikey}
}

func (c *TwitterConnector) Login() string {
	// establish connection with Twitter api
	// conn, err := twitterApi.Connect(c.apikey)
	// c.conn = conn
	return "logged in to twitter"
}

func (c TwitterConnector) LogOut() string {
	// c.conn.Logout()
	return "logged out from twitter"
}

func (c TwitterConnector) CreatePost(post string) string {
	// c.conn.CreatePost(post)
	return fmt.Sprintf("posting [%s] to twitter", post)
}
