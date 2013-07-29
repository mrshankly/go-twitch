package twitch

import (
	"net/url"
	"strconv"
)

type TeamsS struct {
	Teams []TeamS `json:"teams:omitempty"`
	Links LinksS  `json:"_links:omitempty"`
}

type TeamsMethod struct {
	client *Client
}

func (t *TeamsMethod) List(opt *ListOptions) (*TeamsS, error) {
	rel := "teams"
	if opt != nil {
		p := url.Values{
			"limit":  []string{strconv.Itoa(opt.Limit)},
			"offset": []string{strconv.Itoa(opt.Offset)},
		}
		rel += "?" + p.Encode()
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
