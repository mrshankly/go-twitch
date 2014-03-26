package twitch

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

const rootURL = "https://api.twitch.tv/kraken/"

type Client struct {
	client   *http.Client
	BaseURL  *url.URL
	ClientId string

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

	clientId := os.Getenv("GO-TWITCH_CLIENTID")
	c := &Client{client: httpClient, BaseURL: baseURL, ClientId: clientId}
	c.Channels = &ChannelsMethod{client: c}
	c.Chat = &ChatMethod{client: c}
	c.Games = &GamesMethod{client: c}
	c.Ingests = &IngestsMethod{client: c}
	c.Search = &SearchMethod{client: c}
	c.Streams = &StreamsMethod{client: c}
	c.Teams = &TeamsMethod{client: c}
	c.Users = &UsersMethod{client: c}
	c.Videos = &VideosMethod{client: c}

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

	req, err := http.NewRequest("GET", u.String(), nil)

	if err != nil {
		return nil, err

	}
	req.Header.Add("Accept", "application/vnd.twitchtv.v2+json")

	if len(c.ClientId) != 0 {
		req.Header.Add("Client-ID", c.ClientId)

	}
	resp, err := c.client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNotModified {
		return nil, errors.New("api error, response code: " + strconv.Itoa(resp.StatusCode))
	}

	defer resp.Body.Close()

	if r != nil {
		err = json.NewDecoder(resp.Body).Decode(r)
	}

	return resp, err
}
