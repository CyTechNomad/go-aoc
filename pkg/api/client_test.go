package api

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/bilrik/go-aoc/pkg/models"
	"github.com/stretchr/testify/assert"
)

var (
	yearfn = func() int {
		date := now(timezone)
		if date.Month() == time.December {
			return date.Year()
		} else {
			return date.Year() - 1
		}
	}
	dayfn = func() int {
		return now(timezone).Day()

	}
	timezone, _ = time.LoadLocation(models.AOCTimezone)
	date        = time.Now().In(timezone)
)

func Test_NewClient(t *testing.T) {
	tests := map[string]struct {
		opts       []ClientOption
		wantClient *Client
		configfn   func()
	}{
		"No Options": {
			wantClient: &Client{
				Timezone: timezone,
				Year: func() int {
					date := now(timezone)
					if date.Month() == time.December {
						return date.Year()
					} else {
						return date.Year() - 1
					}
				},
				Day: dayfn,
				User: &models.User{
					SessionToken: os.Getenv("AOC_SESSION"),
				},
				HTTPClient: &http.Client{},
			},
		},
		"No Options - !December": {
			wantClient: &Client{
				Timezone: timezone,
				Year:     yearfn,
				Day:      dayfn,
				User: &models.User{
					SessionToken: os.Getenv("AOC_SESSION"),
				},
				HTTPClient: &http.Client{},
			},
			configfn: func() {
				now = func(loc *time.Location) time.Time {
					return time.Date(date.Year(), time.November, date.Day(), 0, 0, 0, 0, loc)
				}
			},
		},
		"With Year": {
			opts: []ClientOption{
				WithYear(2020),
			},
			wantClient: &Client{
				Timezone: timezone,
				Year:     func() int { return 2020 },
				Day:      dayfn,
				User: &models.User{
					SessionToken: os.Getenv("AOC_SESSION"),
				},
				HTTPClient: &http.Client{},
			},
		},
		"With Day": {
			opts: []ClientOption{
				WithDay(1),
			},
			wantClient: &Client{
				Timezone: timezone,
				Year:     yearfn,
				Day:      func() int { return 1 },
				User: &models.User{
					SessionToken: os.Getenv("AOC_SESSION"),
				},
				HTTPClient: &http.Client{},
			},
		},
		"With User": {
			opts: []ClientOption{
				WithUser(&models.User{
					SessionToken: "testToken",
				}),
			},
			wantClient: &Client{
				Timezone: timezone,
				Year:     yearfn,
				Day:      dayfn,
				User: &models.User{
					SessionToken: "testToken",
				},
				HTTPClient: &http.Client{},
			},
		},
		"With Timezone": {
			opts: []ClientOption{
				WithTimezone(time.UTC),
			},
			wantClient: &Client{
				Timezone: time.UTC,
				Year:     yearfn,
				Day:      dayfn,
				User: &models.User{
					SessionToken: os.Getenv("AOC_SESSION"),
				},
				HTTPClient: &http.Client{},
			},
		},
		"With All Options": {
			opts: []ClientOption{
				WithTimezone(time.UTC),
				WithYear(2020),
				WithDay(1),
				WithUser(&models.User{
					SessionToken: "testToken",
				}),
			},
			wantClient: &Client{
				Timezone: time.UTC,
				Year:     func() int { return 2020 },
				Day:      func() int { return 1 },
				User: &models.User{
					SessionToken: "testToken",
				},
				HTTPClient: &http.Client{},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			now = func(loc *time.Location) time.Time {
				return time.Now().In(loc)
			}

			if test.configfn != nil {
				test.configfn()
			}

			gotClient := NewClient(test.opts...)

			assert.ObjectsAreEqualValues(test.wantClient, gotClient)
		})
	}
}
