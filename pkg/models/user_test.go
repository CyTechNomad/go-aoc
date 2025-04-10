package models

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewUser(t *testing.T) {
	tests := map[string]struct {
		opts     []UserOption
		wantUser *User
	}{
		"No_opts": {
			wantUser: &User{
				SessionToken: os.Getenv("AOC_SESSION"),
			},
		},
		"With_session_token": {
			opts: []UserOption{
				WithUserSessionToken("sessionToken"),
			},
			wantUser: &User{
				SessionToken: "sessionToken",
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			gotUser := NewUser(test.opts...)

			assert.EqualValues(t, test.wantUser, gotUser)
		})
	}
}
