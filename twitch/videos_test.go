package twitch

import (
	"net/http"
	"testing"
)

func TestVideosId(t *testing.T) {

	tc := NewClient(&http.Client{})

	_, err := tc.Videos.Id("a328087483")

	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}

}

func TestVideosTop(t *testing.T) {

	tc := NewClient(&http.Client{})

	opt := &ListOptions{
		Limit:  1,
		Offset: 0,
		Game:   "Diablo",
		Period: "week",
	}

	_, err := tc.Videos.Top(opt)

	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}

}
