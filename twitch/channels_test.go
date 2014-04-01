package twitch

import (
	"net/http"
	"testing"
)

func TestChannelsVideos(t *testing.T) {
	tc := NewClient(&http.Client{})

	opt := &ListOptions{
		Limit:  1,
		Offset: 1,
	}

	_, err := tc.Channels.Videos("Dansgaming", opt)

	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}
}

func TestChannelsFollows(t *testing.T) {
	tc := NewClient(&http.Client{})

	opt := &ListOptions{
		Limit:  1,
		Offset: 1,
	}

	_, err := tc.Channels.Follows("Dansgaming", opt)

	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}
}

func TestChannelsChannel(t *testing.T) {
	tc := NewClient(&http.Client{})

	_, err := tc.Channels.Channel("Dansgaming")

	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}
}
