package twitch

import (
	"time"
)

type Channel struct {
	ID                           int64     `json:"_id,omitempty"`
	Mature                       bool      `json:"mature,omitempty"`
	Status                       string    `json:"status,omitempty"`
	BroadcasterLanguage          string    `json:"broadcaster_language,omitempty"`
	DisplayName                  string    `json:"display_name,omitempty"`
	Game                         string    `json:"game,omitempty"`
	Language                     string    `json:"language,omitempty"`
	Name                         string    `json:"name,omitempty"`
	CreatedAt                    time.Time `json:"created_at,omitempty"`
	UpdatedAt                    time.Time `json:"updated_at,omitempty"`
	Partner                      bool      `json:"partner,omitempty"`
	Logo                         string    `json:"logo,omitempty"`
	VideoBanner                  string    `json:"video_banner,omitempty"`
	ProfileBanner                string    `json:"profile_banner,omitempty"`
	ProfileBannerBackgroundColor string    `json:"profile_banner_background_color,omitempty"`
	URL                          string    `json:"url,omitempty"`
	Views                        int       `json:"views,omitempty"`
	Followers                    int       `json:"followers,omitempty"`
	BroadcasterType              string    `json:"broadcaster_type,omitempty"`
	StreamKey                    string    `json:"stream_key,omitempty"`
	Email                        string    `json:"email,omitempty"`
}
