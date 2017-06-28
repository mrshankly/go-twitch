package twitch

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const clipsPath = "clips"

// ClipsService handles communication with the clips related methods of the
// Twitch API.
//
// Twitch API docs: https://dev.twitch.tv/docs/v5/reference/clips/
type ClipsService service

type ClipChannel struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	ChannelURL  string `json:"channel_url,omitempty"`
	Logo        string `json:"logo,omitempty"`
}

type Clip struct {
	Slug        string       `json:"slug,omitempty"`
	TrackingID  string       `json:"tracking_id,omitempty"`
	URL         string       `json:"url,omitempty"`
	EmbedURL    string       `json:"embed_url,omitempty"`
	EmbedHTML   string       `json:"embed_html,omitempty"`
	Broadcaster *ClipChannel `json:"broadcaster,omitempty"`
	Curator     *ClipChannel `json:"curator,omitempty"`
	Vod         *struct {
		ID  string `json:"id,omitempty"`
		URL string `json:"url,omitempty"`
	} `json:"vod,omitempty"`
	Game       string    `json:"game,omitempty"`
	Language   string    `json:"language,omitempty"`
	Title      string    `json:"title,omitempty"`
	Views      int       `json:"views,omitempty"`
	Duration   float64   `json:"duration,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	Thumbnails *struct {
		Medium string `json:"medium,omitempty"`
		Small  string `json:"small,omitempty"`
		Tiny   string `json:"tiny,omitempty"`
	} `json:"thumbnails,omitempty"`
}

type clipsRoot struct {
	Clips  []*Clip `json:"clips,omitempty"`
	Cursor string  `json:"_cursor,omitempty"`
}

type ClipsOptions struct {
	// If this is not specified, top Clips for all channels are returned. If both
	// channel and game are specified, game is ignored. Maximum: 10.
	Channels []string `url:"channel,omitempty,comma"`

	// Game names can be retrieved with the Search Games endpoint. If this is not
	// specified, top Clips for all games are returned. If both channel and game
	// are specified, game is ignored. Maximum: 10.
	Games []string `url:"game,omitempty,comma"`

	// Constrains the languages of videos returned. If no language is specified,
	// all languages are returned. Maximum: 28.
	Languages []string `url:"language,omitempty,comma"`

	// The window of time to search for Clips.
	Period string `url:"period,omitempty"`

	// If true, the Clips returned are ordered by popularity, otherwise, by
	// viewcount.
	Trending bool `url:"trending,omitempty"`

	ListOptions
}

// Returns the object of a specified Clip. Clips are referenced by a globally
// unique string called a slug.
//
// Twitch API docs: https://dev.twitch.tv/docs/v5/reference/clips/#get-clip
func (s *ClipsService) Clip(ctx context.Context, slug string) (*Clip, *http.Response, error) {
	path := fmt.Sprintf("%v/%v", clipsPath, slug)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	clip := new(Clip)
	resp, err := s.client.Do(ctx, req, clip)
	if err != nil {
		return nil, resp, err
	}

	return clip, resp, nil
}

// Returns the top Clips which meet a specified set of parameters.
//
// Twitch API docs: https://dev.twitch.tv/docs/v5/reference/clips/#get-top-clips
func (s *ClipsService) Top(ctx context.Context, opt *ClipsOptions) ([]*Clip, string, *http.Response, error) {
	root := new(clipsRoot)
	resp, err := s.list(ctx, "top", opt, root)
	return root.Clips, root.Cursor, resp, err
}

// Returns the top Clips for the games followed by the authenticated user.
//
// Twitch API docs: https://dev.twitch.tv/docs/v5/reference/clips/#get-followed-clips
func (s *ClipsService) Followed(ctx context.Context, opt *ClipsOptions) ([]*Clip, string, *http.Response, error) {
	root := new(clipsRoot)
	resp, err := s.list(ctx, "followed", opt, root)
	return root.Clips, root.Cursor, resp, err
}

func (s *ClipsService) list(ctx context.Context, clipsType string, opt *ClipsOptions, r interface{}) (*http.Response, error) {
	path := fmt.Sprintf("%v/%v", clipsPath, clipsType)
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, r)
}
