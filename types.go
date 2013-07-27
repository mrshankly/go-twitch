package twitch

type StreamS struct {
	Preview *PreviewS `json:"preview,omitempty"`
	Viewers int       `json:"viewers,omitempty"`
	Game    string    `json:"game,omitempty"`
}

type FStreamS struct {
	Stream *StreamS `json:"stream,omitempty"`
	Text   string   `json:"text,omitempty"`
	Image  string   `json:"image,omitempty"`
}

type LinksS struct {
	Summary  string `json:"summary,omitempty"`
	Followed string `json:"followed,omitempty"`
	Next     string `json:"streams,omitempty"`
	Featured string `json:"featured,omitempty"`
}

type PreviewS struct {
	Small    string `json:"small,omitempty"`
	Medium   string `json:"medium,omitempty"`
	Large    string `json:"large,omitempty"`
	Template string `json:"template,omitempty"`
}
