package twitch

import (
	"net/http"
	"testing"
)

func TestStreamsChannel(t *testing.T) {
	tc := NewClient(&http.Client{})

	_, err := tc.Streams.Channel("Robbaz")

	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}
}

func TestStreamsList(t *testing.T) {
	tc := NewClient(&http.Client{})

	emb := true
	hls := false

	opt := &ListOptions{
		Game:       "DayZ",
		Channel:    "LIRIK",
		Limit:      1,
		Offset:     1,
		Embeddable: &emb,
		Hls:        &hls,
	}

	_, err := tc.Streams.List(opt)

	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}
}

func TestStreamsFeatured(t *testing.T) {
	tc := NewClient(&http.Client{})

	opt := &ListOptions{
		Limit:  1,
		Offset: 1,
	}

	_, err := tc.Streams.Featured(opt)

	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}
}

func TestStreamsSummary(t *testing.T) {
	tc := NewClient(&http.Client{})

	opt := &ListOptions{
		Game: "DayZ",
	}

	_, err := tc.Streams.Summary(opt)

	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}
}
