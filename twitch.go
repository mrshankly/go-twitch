package twitch

import (
	"net/http"
	"net/url"
	"encoding/json"
)

const rootURL   = "https://api.twitch.tv/kraken/"

type Client struct {
	client *http.Client
	BaseURL *url.URL

	// Twitch api methods
	Streams *StreamsMethod
}

// Returns a new twitch client used to communicate with the API.
func NewClient(httpClient *http.Client) *Client {
	baseURL, _ := url.Parse(rootURL)

	c := &Client{client: httpClient, BaseURL: baseURL}
	c.Streams = &StreamsMethod{client: c}

	return c
}

// Issues an API get request and returns the API response. The response body is
// decoded and stored in the value pointed by r.
func (c *Client) Get(path string, r interface{}) (*http.Response, error) {
	rel, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	resp, err := c.client.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if r != nil {
		err = json.NewDecoder(resp.Body).Decode(r)
	}
	return resp, err
}
