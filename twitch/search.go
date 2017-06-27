package twitch

import (
	"context"
	"fmt"
	"net/http"

	qs "github.com/google/go-querystring/query"
)

const searchPath = "search"

// SearchService handles communication with the search related methods of the
// Twitch API.
//
// Twitch API docs: https://dev.twitch.tv/docs/v5/reference/search/
type SearchService service

type SearchOptions struct {
	// Search only games that are live on at least one channel.
	Live bool `url:"live,omitempty"`

	// If true, search only HLS streams. If false, only RTMP streams. If not set,
	// both HLS and RTMP streams.
	HLS bool `url:"hls,omitempty"`

	ListOptions
}

type searchChannelsRoot struct {
	Channels []*Channel `json:"channels,omitempty"`
}

type searchGamesRoot struct {
	Games []*Game `json:"games,omitempty"`
}

type searchStreamsRoot struct {
	Streams []*Stream `json:"streams,omitempty"`
}

// Searches for channels based on a specified query parameter. A channel is
// returned if the query parameter is matched entirely or partially, in the
// channel description or game name.
//
// Twitch API docs: https://dev.twitch.tv/docs/v5/reference/search/#search-channels
func (s *SearchService) Channels(ctx context.Context, query string, opt *SearchOptions) ([]*Channel, *http.Response, error) {
	root := new(searchChannelsRoot)
	resp, err := s.search(ctx, "channels", query, opt, root)
	return root.Channels, resp, err
}

// Searches for games based on a specified query parameter. A game is returned
// if the query parameter is matched entirely or partially, in the game name.
//
// Twitch API docs: https://dev.twitch.tv/docs/v5/reference/search/#search-games
func (s *SearchService) Games(ctx context.Context, query string, opt *SearchOptions) ([]*Game, *http.Response, error) {
	root := new(searchGamesRoot)
	resp, err := s.search(ctx, "games", query, opt, root)
	return root.Games, resp, err
}

// Searches for streams based on a specified query parameter. A stream is
// returned if the query parameter is matched entirely or partially, in the
// channel description or game name.
//
// Twitch API docs: https://dev.twitch.tv/docs/v5/reference/search/#search-streams
func (s *SearchService) Streams(ctx context.Context, query string, opt *SearchOptions) ([]*Stream, *http.Response, error) {
	root := new(searchStreamsRoot)
	resp, err := s.search(ctx, "streams", query, opt, root)
	return root.Streams, resp, err
}

func (s *SearchService) search(ctx context.Context, searchType, query string, opt *SearchOptions, r interface{}) (*http.Response, error) {
	params, err := qs.Values(opt)
	if err != nil {
		return nil, err
	}

	params.Set("query", query)
	path := fmt.Sprintf("%s/%s?%s", searchPath, searchType, params.Encode())

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, r)
}
