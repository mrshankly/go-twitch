package twitch

import (
	"time"
)

type Stream struct {
	ID          int64     `json:"_id,omitempty"`
	Game        string    `json:"game,omitempty"`
	Viewers     int       `json:"viewers,omitempty"`
	VideoHeight int       `json:"video_height,omitempty"`
	AverageFps  float64   `json:"average_fps,omitempty"`
	Delay       int       `json:"delay,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	IsPlaylist  bool      `json:"is_playlist,omitempty"`
	Preview     *Preview  `json:"preview,omitempty"`
	Channel     *Channel  `json:"channel,omitempty"`
}
