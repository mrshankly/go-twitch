package twitch

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const teamsPath = "teams"

// TeamsService handles communication with the teams related methods of the
// Twitch API.
//
// Twitch API docs: https://dev.twitch.tv/docs/v5/reference/teams/
type TeamsService service

type User struct {
	ID                           string    `json:"_id,omitempty"`
	BroadcasterLanguage          string    `json:"broadcaster_language,omitempty"`
	CreatedAt                    time.Time `json:"created_at,omitempty"`
	DisplayName                  string    `json:"display_name,omitempty"`
	Followers                    int       `json:"followers,omitempty"`
	Game                         string    `json:"game,omitempty"`
	Language                     string    `json:"language,omitempty"`
	Logo                         string    `json:"logo,omitempty"`
	Mature                       bool      `json:"mature,omitempty"`
	Name                         string    `json:"name,omitempty"`
	Partner                      bool      `json:"partner,omitempty"`
	ProfileBanner                string    `json:"profile_banner,omitempty"`
	ProfileBannerBackgroundColor string    `json:"profile_banner_background_color,omitempty"`
	Status                       string    `json:"status,omitempty"`
	UpdatedAt                    time.Time `json:"updated_at,omitempty"`
	URL                          string    `json:"url,omitempty"`
	VideoBanner                  string    `json:"video_banner,omitempty"`
	Views                        int       `json:"views,omitempty"`
}

type Team struct {
	ID          int       `json:"_id,omitempty"`
	Background  string    `json:"background,omitempty"`
	Banner      string    `json:"banner,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	DisplayName string    `json:"display_name,omitempty"`
	Info        string    `json:"info,omitempty"`
	Logo        string    `json:"logo,omitempty"`
	Name        string    `json:"name,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	Users       []*User   `json:"users,omitempty"`
}

type listTeamsRoot struct {
	Teams []*Team `json:"teams,omitempty"`
}

// Returns all active teams.
//
// Twitch API docs: https://dev.twitch.tv/docs/v5/reference/teams/#get-all-teams
func (s *TeamsService) List(ctx context.Context, opt *ListOptions) ([]*Team, *http.Response, error) {
	path, err := addOptions(teamsPath, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(listTeamsRoot)
	resp, err := s.client.Do(ctx, req, &root)
	if err != nil {
		return nil, resp, err
	}

	return root.Teams, resp, nil
}

// Returns a specified team object.
//
// Twitch API docs: https://dev.twitch.tv/docs/v5/reference/teams/#get-team
func (s *TeamsService) Team(ctx context.Context, name string) (*Team, *http.Response, error) {
	path := fmt.Sprintf("%v/%v", teamsPath, name)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	team := new(Team)
	resp, err := s.client.Do(ctx, req, &team)
	if err != nil {
		return nil, resp, err
	}

	return team, resp, nil
}
