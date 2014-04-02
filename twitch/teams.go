package twitch

import "github.com/google/go-querystring/query"

type TeamsS struct {
	Teams []TeamS `json:"teams,omitempty"`
	Links LinksS  `json:"_links,omitempty"`
}

type TeamsMethod struct {
	client *Client
}

func (t *TeamsMethod) List(opt *ListOptions) (*TeamsS, error) {
	rel := "teams"
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "?" + v.Encode()
	}

	teams := new(TeamsS)
	_, err := t.client.Get(rel, teams)
	return teams, err
}

func (t *TeamsMethod) Team(name string) (*TeamS, error) {
	rel := "teams/" + name

	team := new(TeamS)
	_, err := t.client.Get(rel, team)
	return team, err
}
