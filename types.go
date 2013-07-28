package twitch

type ChannelS struct {
	Name        string         `json:"name,omitempty"`
	Game        string         `json:"game,omitempty"`
	Delay       int            `json:"game,omitempty"`
	Teams       []*TeamS       `json:"teams,omitempty"`
	Title       string         `json:"title,omitempty"`
	Banner      string         `json:"banner,omitempty"`
	VideoBanner string         `json:"video_banner,omitempty"`
	Background  string         `json:"background,omitempty"`
	Links       *ChannelLinksS `json:"_links,omitempty"`
	Logo        string         `json:"logo,omitempty"`
	Mature      bool           `json:"mature,omitempty"`
	Url         string         `json:"url,omitempty"`
	DisplayName string         `json:"display_name,omitempty"`
}

type TeamS struct {
	Name        string `json:"name,omitempty"`
	Background  string `json:"background,omitempty"`
	Banner      string `json:"banner,omitempty"`
	Logo        string `json:"logo,omitempty"`
	Info        string `json:"info,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
}

type ChannelLinksS struct {
    Chat       string `json:"chat,omitempty"`
    Commercial string `json:"commercial,omitempty"`
    Videos     string `json:"videos,omitempty"`
}

type StreamS struct {
	Game        string    `json:"game,omitempty"`
	Name        string    `json:"name,omitempty"`
	Preview     *PreviewS `json:"preview,omitempty"`
	Viewers     int       `json:"viewers,omitempty"`
	Broadcaster string    `json:"broadcaster,omitempty"`
	Geo         string    `json:"geo,omitempty"`
	Channel     *ChannelS `json:"channel,omitempty"`
	Status      string    `json:"status,omitempty"`
}

type FStreamS struct {
	Stream *StreamS `json:"stream,omitempty"`
	Text   string   `json:"text,omitempty"`
	Image  string   `json:"image,omitempty"`
}

type VideoS struct {
	Title       string `json:"title,omitempty"`
	ID          string `json:"_id,omitempty"`
	Embed       string `json:"embed,omitempty"`
	Url         string `json:"url,omitempty"`
	Views       int    `json:"views,omitempty"`
	Preview     string `json:"preview,omitempty"`
	Length      int    `json:"length,omitempty"`
	Description string `json:"description,omitempty"`
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

type ListOptions struct {
	Game       string
	Channel    string
	Limit      int
	Offset     int
	Embeddable bool
	Hls        bool
	ClientId   string
}
