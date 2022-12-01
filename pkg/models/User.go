package models

import (
	"net/http"
	"net/url"
	"os"
)

// User is the user data.
type User struct {
	sessionToken string
}

// NewUser creates a new user.
func NewUser() *User {
	return &User{
		sessionToken: os.Getenv("AOC_SESSION"),
	}
}

// GetUserSessionToken returns the user sessionToken.
func (u *User) GetUserSessionToken() string {
	return u.sessionToken
}

// SetUserSessionToken sets the user sessionToken.
func (u *User) SetUserSessionToken(sessionToken string) {
	u.sessionToken = sessionToken
}

// GetHeaders returns the headers.
func (u *User) GetHeaders() http.Header {
	headers := make(http.Header)
	headers.Set("Cookie", "session="+u.sessionToken)
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
