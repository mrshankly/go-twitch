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
		p := url.Values{
			"limit":  []string{strconv.Itoa(opt.Limit)},
			"offset": []string{strconv.Itoa(opt.Offset)},
		}
		rel += "&" + p.Encode()
	}

	search := new(StreamsS)
	_, err := s.client.Get(rel, search)
	return search, err
}

func (s *SearchMethod) Games(q string, opt *ListOptions) (*SGamesS, error) {
	rel := fmt.Sprintf("search/games?q=%s&type=suggest", url.QueryEscape(q))
	if opt != nil {
		p := url.Values{"live": []string{strconv.FormatBool(opt.Live)}}
		rel += "&" + p.Encode()
	}

	search := new(SGamesS)
	_, err := s.client.Get(rel, search)
	return search, err
}
