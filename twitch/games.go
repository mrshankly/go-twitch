package twitch

import (
	"net/url"
	"strconv"
)

type TopsS struct {
	Links LinksS `json:"_links,omitempty"`
	Total int    `json:"_total,omitempty"`
	Top   []TopS `json:"top,omitempty"`
}

type TopS struct {
	Game     GameS `json:"game,omitempty"`
	Viewers  int   `json:"viewers,omitempty"`
	Channels int   `json:"channels,omitempty"`
}

type GameS struct {
	Name        string   `json:"name,omitempty"`
	Box         PreviewS `json:"box,omitempty"`
	Logo        PreviewS `json:"logo,omitempty"`
	Id          int      `json:"_id,omitempty"`
	GiantbombId int      `json:"giantbomb_id,omitempty"`
}

type GamesMethod struct {
	client *Client
}

// Returns a list of games objects sorted by number of current viewers, most
// popular first.
func (g *GamesMethod) Top(opt *ListOptions) (*TopsS, error) {
	rel := "games/top"
	if opt != nil {
		p := url.Values{}
		if opt.Limit > 0 {
			p.Add("limit", strconv.Itoa(opt.Limit))
		}
		if opt.Offset > 0 {
			p.Add("offset", strconv.Itoa(opt.Offset))
		}
		if opt.Hls != nil {
			p.Add("hls", strconv.FormatBool(opt.Hls.Show))
		}
		if len(p) > 0 {
			rel += "?" + p.Encode()
		}
	}

	games := new(TopsS)
	_, err := g.client.Get(rel, games)
	return games, err
}
