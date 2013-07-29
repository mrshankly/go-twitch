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
	Channels *ChannelsMethod
	Chat     *ChatMethod
	Games    *GamesMethod
	Ingests  *IngestsMethod
	Search   *SearchMethod
	Streams  *StreamsMethod
	Teams    *TeamsMethod
	Users    *UsersMethod
	Videos   *VideosMethod
}

// Returns a new twitch client used to communicate with the API.
func NewClient(httpClient *http.Client) *Client {
	baseURL, _ := url.Parse(rootURL)

	c := &Client{client: httpClient, BaseURL: baseURL}
	c.Channels = &ChannelsMethod{client: c}
	c.Chat     = &ChatMethod{client: c}
	c.Games    = &GamesMethod{client: c}
	c.Ingests  = &IngestsMethod{client: c}
	c.Search   = &SearchMethod{client: c}
	c.Streams  = &StreamsMethod{client: c}
	c.Teams    = &TeamsMethod{client: c}
	c.Users    = &UsersMethod{client: c}
	c.Videos   = &VideosMethod{client: c}

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
