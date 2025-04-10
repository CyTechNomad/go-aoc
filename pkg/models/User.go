package models

import (
	"net/http"
	"net/url"
	"os"
)

// User is the user data.
type User struct {
	SessionToken string
}

type UserOption func(*User)

// NewUser creates a new user.
func NewUser(opts ...UserOption) *User {
	// create empty user
	u := &User{}

	// applies provided options
	for _, opt := range opts {
		opt(u)
	}

	// default options
	// // may need to be moved to run before opts if more options are added
	// // so that opts can override defaults if needed
	if u.SessionToken == "" {
		defaultOptions(u)
	}

	return u
}

// defaultOptions sets the default options.
// // default sessionToken: is attempted to be set from environment
func defaultOptions(u *User) {
	u.SessionToken = os.Getenv("AOC_SESSION")
	if u.SessionToken == "" {
		panic("AOC_SESSION environment variable not set and none provided")
	}
}

// withSessionToken sets the sessionToken.
func WithUserSessionToken(sessionToken string) UserOption {
	return func(u *User) {
		u.SessionToken = sessionToken
	}
}

// GetUserSessionToken returns the user sessionToken.
func (u *User) GetUserSessionToken() string {
	return u.SessionToken
}

// GetHeaders returns the headers.
func (u *User) GetHeaders() http.Header {
	headers := make(http.Header)
	headers.Set("Cookie", "session="+u.SessionToken)
	headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:107.0) Gecko/20100101 Firefox/107.0")

	return headers
}

// SetFormValues sets the form values.
func (u *User) SetFormValues(req *http.Request, part, answer string) url.Values {
	req.Header = u.GetHeaders()
	req.Form = make(url.Values)
	req.Form.Set("level", part)
	req.Form.Set("answer", answer)

	return req.Form
}
