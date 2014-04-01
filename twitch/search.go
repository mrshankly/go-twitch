package twitch

import (
	"fmt"
	"net/url"
	"strconv"
)

type SGamesS struct {
	Games []SGameS `json:"games,omitempty"`
}

type SGameS struct {
	Box         PreviewS `json:"box,omitempty"`
	Logo        PreviewS `json:"logo,omitempty"`
	Images      ImagesS  `json:"images,omitempty"`
	Popularity  int      `json:"popularity,omitempty"`
	Name        string   `json:"name,omitempty"`
	Id          int      `json:"_id,omitempty"`
	GiantbombId int      `json:"giantbomb_id,omitempty"`
}

type ImagesS struct {
	Thumb  string `json:"thumb,omitempty"`
	Tiny   string `json:"tiny,omitempty"`
	Small  string `json:"small,omitempty"`
	Super  string `json:"super,omitempty"`
	Medium string `json:"medium,omitempty"`
	Icon   string `json:"icon,omitempty"`
	Screen string `json:"screen,omitempty"`
}

type SearchMethod struct {
	client *Client
}

func (s *SearchMethod) Streams(q string, opt *ListOptions) (*StreamsS, error) {
	rel := "search/streams?q=" + url.QueryEscape(q)
	if opt != nil {
		p := url.Values{}
		if opt.Limit > 0 {
			p.Add("limit", strconv.Itoa(opt.Limit))
		}
		if opt.Offset > 0 {
			p.Add("offset", strconv.Itoa(opt.Offset))
		}
		if len(p) > 0 {
			rel += "&" + p.Encode()
		}
	}

	search := new(StreamsS)
	_, err := s.client.Get(rel, search)
	return search, err
}

func (s *SearchMethod) Games(q string, opt *ListOptions) (*SGamesS, error) {
	rel := fmt.Sprintf("search/games?q=%s&type=suggest", url.QueryEscape(q))
	if opt != nil {
		p := url.Values{}
		if opt.Live != nil {
			p.Add("live", strconv.FormatBool(opt.Live.Show))
		}
		if len(p) > 0 {
			rel += "&" + p.Encode()
		}
	}

	search := new(SGamesS)
	_, err := s.client.Get(rel, search)
	return search, err
}
