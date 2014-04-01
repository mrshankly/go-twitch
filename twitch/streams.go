// Streams methods of the twitch api.
// https://github.com/justintv/Twitch-API/blob/master/v3_resources/streams.md

package twitch

import (
	"net/url"
	"strconv"
)

// used with GET /streams/:channel/
type SChannelS struct {
	Stream StreamS `json:"stream,omitempty"`
	Links  LinksS  `json:"_links,omitempty"`
}

// used with GET /streams
type StreamsS struct {
	Total   int       `json:"_total,omitempty"`
	Streams []StreamS `json:"streams,omitempty"`
	Links   LinksS    `json:"_links,omitempty"`
}

// used with GET /streams/featured
type FeaturedS struct {
	Featured []FStreamS `json:"featured,omitempty"`
	Links    LinksS     `json:"_links,omitempty"`
}

// used with GET /streams/summary
type SummaryS struct {
	Viewers  int `json:"viewers,omitempty"`
	Channels int `json:"channels,omitempty"`
}

// used with GET /streams/followed
type FollowedS struct {
	Streams []StreamS `json:"streams,omitempty"`
	Links   LinksS    `json:"_links,omitempty"`
}

type StreamsMethod struct {
	client *Client
}

// Returns a stream object if live.
func (s *StreamsMethod) Channel(name string) (*SChannelS, error) {
	rel := "streams/" + name

	stream := new(SChannelS)
	_, err := s.client.Get(rel, stream)
	return stream, err
}

// Returns a list of stream objects according to optional parameters.
func (s *StreamsMethod) List(opt *ListOptions) (*StreamsS, error) {
	rel := "streams"

	if opt != nil {
		p := url.Values{}
		if len(opt.Game) > 0 {
			p.Add("game", opt.Game)
		}
		if len(opt.Channel) > 0 {
			p.Add("channel", opt.Channel)
		}
		if opt.Limit > 0 {
			p.Add("limit", strconv.Itoa(opt.Limit))
		}
		if opt.Offset > 0 {
			p.Add("offset", strconv.Itoa(opt.Offset))
		}
		if opt.Hls != nil {
			p.Add("hls", strconv.FormatBool(opt.Hls.Show))
		}
		if opt.Embeddable != nil {
			p.Add("embeddable", strconv.FormatBool(opt.Embeddable.Show))
		}
		if len(p) > 0 {
			rel += "?" + p.Encode()
		}
	}

	streams := new(StreamsS)
	_, err := s.client.Get(rel, streams)
	return streams, err
}

// Returns a list of featured (promoted) stream objects.
func (s *StreamsMethod) Featured(opt *ListOptions) (*FeaturedS, error) {
	rel := "streams/featured"
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

	featured := new(FeaturedS)
	_, err := s.client.Get(rel, featured)
	return featured, err
}

// Returns a summary of current streams.
func (s *StreamsMethod) Summary(opt *ListOptions) (*SummaryS, error) {
	rel := "streams/summary"
	if opt != nil {
		p := url.Values{}
		if len(opt.Game) > 0 {
			p.Add("game", opt.Game)
		}
		if len(p) > 0 {
			rel += "?" + p.Encode()
		}
	}

	summary := new(SummaryS)
	_, err := s.client.Get(rel, summary)
	return summary, err
}

// Returns a list of stream objects that the authenticated user is following.
func (s *StreamsMethod) followed(opt *ListOptions) (*FollowedS, error) {
	rel := "streams/followed"
	if opt != nil {
		p := url.Values{
			"limit":  []string{strconv.Itoa(opt.Limit)},
			"offset": []string{strconv.Itoa(opt.Offset)},
			//"hls":    []string{strconv.FormatBool(opt.Hls)},
		}
		rel += "?" + p.Encode()
	}

	followed := new(FollowedS)
	_, err := s.client.Get(rel, followed)
	return followed, err
}
