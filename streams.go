// Streams methods of the twitch api.
// https://github.com/justintv/Twitch-API/blob/master/v3_resources/streams.md

package twitch

import (
	"strconv"
	"net/url"
)

// used with GET /streams/:channel/
type SChannelS struct {
	Stream *StreamS `json:"stream,omitempty"`
	Links  *LinksS  `json:"_links,omitempty"`
}

// used with GET /streams
type StreamsS struct {
	Streams []*StreamS `json:"streams,omitempty"`
	Links   *LinksS    `json:"_links,omitempty"`
}

// used with GET /streams/featured
type FeaturedS struct {
	Featured []*FStreamS `json:"featured,omitempty"`
	Links    *LinksS     `json:"_links,omitempty"`
}

// used with GET /streams/summary
type SummaryS struct {
	Viewers  int `json:"viewers,omitempty"`
	Channels int `json:"channels,omitempty"`
}

// Returns a stream object if online.
func (s *Method) Channel(name string) (*SChannelS, error) {
	rel := "streams/" + name

	stream := new(SChannelS)
	_, err := s.client.Get(rel, stream)
	return stream, err
}

// Returns a list of stream objects according to optional parameters.
func (s *Method) List(opt *ListOptions) (*StreamsS, error) {
	rel := "streams"
	if opt != nil {
		p := url.Values{
			"game":       []string{opt.Game},
			"channel":    []string{opt.Channel},
			"limit":      []string{strconv.Itoa(opt.Limit)},
			"offset":     []string{strconv.Itoa(opt.Offset)},
			"embeddable": []string{strconv.FormatBool(opt.Embeddable)},
			"hls":        []string{strconv.FormatBool(opt.Hls)},
			"client_id":  []string{opt.ClientId},
		}
		rel += "?" + p.Encode()
	}

	streams := new(StreamsS)
	_, err := s.client.Get(rel, streams)
	return streams, err
}

// Returns a list of featured (promoted) stream objects.
func (s *Method) Featured(opt *ListOptions) (*FeaturedS, error) {
	rel := "streams/featured"
	if opt != nil {
		p := url.Values{
			"limit":  []string{strconv.Itoa(opt.Limit)},
			"offset": []string{strconv.Itoa(opt.Offset)},
			"hls":    []string{strconv.FormatBool(opt.Hls)},
		}
		rel += "?" + p.Encode()
	}

	featured := new(FeaturedS)
	_, err := s.client.Get(rel, featured)
	return featured, err
}

// Returns a summary of current streams.
func (s *Method) Summary(opt *ListOptions) (*SummaryS, error) {
	rel := "streams/summary"
	if opt != nil {
		p := url.Values{"game":  []string{opt.Game}}
		rel += "?" + p.Encode()
	}

	summary := new(SummaryS)
	_, err := s.client.Get(rel, summary)
	return summary, err
}

// TODO GET /streams/followed
