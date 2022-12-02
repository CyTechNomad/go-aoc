package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/bilrik/go-aoc/pkg/models"
)

// Client for the API.
// client needs user data and year, day, part
type Client struct {
	Year       int
	Day        int
	User       models.User
	HTTPClient *http.Client
}

// NewClient creates a new client.
func NewClient() *Client {
	var client Client

	date := time.Now().UTC().Add(time.Hour * -5)

	if date.Month() == time.December {
		client.Year = date.Year()
	} else {
		client.Year = date.Year() - 1
	}

	client.Day = date.Day()
	client.User = *models.NewUser()
	client.HTTPClient = &http.Client{}

	return &client
}

// client makes http request to AOC API
func (c *Client) makeRequest(req *http.Request) (*http.Response, error) {

	req.Header = c.User.GetHeaders()
	req.Form = c.User.SetFormValues(req, "2", "5")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// clinet needs to be able to get the data for the day and part
func (c *Client) GetInputData() (*string, error) {
	url := fmt.Sprintf(models.DayURL, c.Year, c.Day) + "/input"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("makeRequest: could not create request: %v", err)
	}

	resp, err := c.makeRequest(req)
	if err != nil {
		return nil, err
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	input := string(respBody)
	return &input, nil
}

// client needs to be able to post the answer for the day and part
func (c *Client) PostAnswer(level, answer string) (*string, error) {
	url := fmt.Sprintf(models.DayURL, c.Year-1, 5) + "/answer"

	//request body need to contain level and answer

	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return nil, fmt.Errorf("makeRequest: could not create request: %v", err)
	}

	resp, err := c.makeRequest(req)
	if err != nil {
		return nil, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}

	returnValue := string(respBody)
	return &returnValue, nil
}
