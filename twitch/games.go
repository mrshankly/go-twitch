package twitch

import (
	"context"
	"fmt"
	"net/http"
)

const gamesPath = "games"

// GamesService handles communication with the games related methods of the
// Twitch API.
//
// Twitch API docs: https://dev.twitch.tv/docs/v5/reference/games/
type GamesService service

type Preview struct {
	Large    string `json:"large,omitempty"`
	Medium   string `json:"medium,omitempty"`
	Small    string `json:"small,omitempty"`
	Template string `json:"template,omitempty"`
}

type Game struct {
	ID          int64    `json:"_id,omitempty"`
	Box         *Preview `json:"box,omitempty"`
	GiantbombID int      `json:"giantbomb_id,omitempty"`
	Logo        *Preview `json:"logo,omitempty"`
	Name        string   `json:"name,omitempty"`
	Popularity  int      `json:"popularity,omitempty"`
}

type TopGame struct {
	Channels int   `json:"channels,omitempty"`
	Viewers  int   `json:"viewers,omitempty"`
	Game     *Game `json:"game,omitempty"`
}

type topGamesRoot struct {
	Top []*TopGame `json:"top,omitempty"`
}

// Returns the games sorted by number of current viewers on Twitch, most popular
// first.
//
// Twitch API docs: https://dev.twitch.tv/docs/v5/reference/games/#get-top-games
func (s *GamesService) Top(ctx context.Context, opt *ListOptions) ([]*TopGame, *http.Response, error) {
	path := fmt.Sprintf("%v/top", gamesPath)
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(topGamesRoot)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.Top, resp, nil
}
