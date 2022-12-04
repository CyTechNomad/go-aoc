package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/bilrik/go-aoc/pkg/models"
)

var now = func(loc *time.Location) time.Time {
	return time.Now().In(loc)
}

type Client struct {
	Timezone   *time.Location
	Year       func() int
	Day        func() int
	User       *models.User
	HTTPClient *http.Client
}

type ClientOption func(*Client)

// NewClient creates a new client.
func NewClient(opts ...ClientOption) *Client {
	// create empty client
	c := &Client{}

	// applies default options
	defaultOptions(c)

	// applies provided options
	for _, opt := range opts {
		opt(c)
	}

	// if user is not provided then attempt to set new default user
	// // this was moved out of defaultOptions so that defaultOptions
	// // does not attempt to create user by default method if a user is provided
	if c.User == nil {
		WithUser(models.NewUser())(c)
	}

	return c
}

func defaultOptions(c *Client) {
	// set default timezone to America/New_York using LoadLocation and check for error
	loc, err := time.LoadLocation(models.AOCTimezone)
	if err != nil {
		panic(err)
	}
	c.Timezone = loc

	c.Year = func() int {
		// set yearfn to return currentyear but only if it is December else set to return previous year
		date := now(c.Timezone)
		if date.Month() == time.December {
			return date.Year()
		} else {
			return date.Year() - 1
		}
	}

	// set day
	c.Day = func() int {
		return now(c.Timezone).Day()
	}

	// create http client
	c.HTTPClient = &http.Client{}
}

// withYear sets the year for the client
func WithYear(year int) ClientOption {
	return func(c *Client) {
		// overide default year
		c.Year = func() int {
			return year
		}
	}
}

// withDay sets the day for the client
func WithDay(day int) ClientOption {
	return func(c *Client) {
		// overide default day
		c.Day = func() int {
			return day
		}
	}
}

// withUser sets the user for the client
func WithUser(user *models.User) ClientOption {
	return func(c *Client) {
		// set user
		c.User = user
	}
}

// withTimezone sets the timezone for the client
func WithTimezone(loc *time.Location) ClientOption {
	return func(c *Client) {
		// overide default timezone
		c.Timezone = loc
	}
}

// client makes http request to AOC API
func (c *Client) makeRequest(req *http.Request) (*http.Response, error) {
	req.Header = c.User.GetHeaders()

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// clinet needs to be able to get the data for the day and part
func (c *Client) GetInputData() (*string, error) {
	// set url with clinet year and day
	url := fmt.Sprintf(models.InputURL, c.Year, c.Day)

	// create request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// make request
	resp, err := c.makeRequest(req)
	if err != nil {
		return nil, err
	}

	// read response body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// return response body as string
	input := string(respBody)

	// return input data
	return &input, nil
}

// client needs to be able to post the answer for the day and level
func (c *Client) PostAnswer(level, answer string) (*string, error) {
	// set url with clinet year and day
	url := fmt.Sprintf(models.AnswerURL, c.Year, c.Day)

	// create request
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}

	// set form values
	req.Form = c.User.SetFormValues(req, level, answer)

	// make request
	resp, err := c.makeRequest(req)
	if err != nil {
		return nil, err
	}

	// read response body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// return response body as string
	returnValue := string(respBody)

	// return input data
	return &returnValue, nil
}
