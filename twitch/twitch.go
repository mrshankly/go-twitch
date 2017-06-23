package twitch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	libraryVersion = "2"
	rootURL        = "https://api.twitch.tv/kraken/"
	userAgent      = "go-twitch/" + libraryVersion
	mediaType      = "application/vnd.twitchtv.v5+json"
)

// A Client manages communication with the Twitch API.
type Client struct {
	client *http.Client

	// Base URL for API requests.
	BaseURL *url.URL

	// User agent used when cummunicating with the Twitch API.
	UserAgent string

	// Twitch client ID.
	ClientID string

	// Services used for talking to different parts of the Twitch API.
	Ingests *IngestsService

	common service
}

type service struct {
	client *Client
}

// Returns a new Twitch API client.
//
// If a nil httpClient is provided, http.DefaultClient will be used. To use API
// methods which require authentication, either set the Twitch client ID with
// SetClientID or provide an http.Client that will perform the authentication
// for you (such as that provided by the golang.org/x/oauth2 library).
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(rootURL)

	c := &Client{
		client:    httpClient,
		BaseURL:   baseURL,
		UserAgent: userAgent,
		ClientID:  "",
	}

	c.common.client = c
	c.Ingests = (*IngestsService)(&c.common)

	return c
}

// Sets the Twitch client ID that will be used by this client.
func (c *Client) SetClientID(id string) error {
	c.ClientID = id
	return nil
}

// Creates an API request.
//
// The path string is resolved relative to the BaseURL of the client.
//
// If not nil, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(method, path string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(path)
	if err != nil {
		return nil, err
	}
	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err = json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if c.ClientID != "" {
		req.Header.Set("Client-ID", c.ClientID)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", mediaType)
	req.Header.Set("User-Agent", c.UserAgent)

	return req, nil
}

// Do sends an API request and returns the API response.
//
// The API response is JSON decoded and stored in the value pointed to by r, or
// returned as an error if an API error has occurred.
//
// The provided ctx must not be nil. If it is canceled or times out, ctx.Err()
// will be returned.
func (c *Client) Do(ctx context.Context, req *http.Request, r interface{}) (*http.Response, error) {
	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		return nil, err
	}
	defer resp.Body.Close()

	if err = checkResponse(resp); err != nil {
		return resp, err
	}

	if r != nil {
		err = json.NewDecoder(resp.Body).Decode(r)
		if err == io.EOF {
			err = nil
		}
	}
	return resp, err
}

// An ErrorResponse reports an error caused by an API request.
type ErrorResponse struct {
	// HTTP response that cause this error.
	Response *http.Response

	// Error message.
	Message string `json:"message,omitempty"`
}

func checkResponse(r *http.Response) error {
	if 200 <= r.StatusCode && r.StatusCode <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && len(data) > 0 {
		err = json.Unmarshal(data, errorResponse)
	}
	return errorResponse
}

func (e *ErrorResponse) Error() string {
	r := e.Response

	return fmt.Sprintf("%v %v: %d %v",
		r.Request.Method, r.Request.URL, r.StatusCode, e.Message)
}
