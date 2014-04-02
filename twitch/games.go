package twitch

import "github.com/google/go-querystring/query"

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
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "?" + v.Encode()
	}

	games := new(TopsS)
	_, err := g.client.Get(rel, games)
	return games, err
}
