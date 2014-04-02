package twitch

import (
	"net/http"
	"testing"
)

func TestSearchStreams(t *testing.T) {
	tc := NewClient(&http.Client{})

	opt := &ListOptions{
		Limit:  1,
		Offset: 0,
	}

	_, err := tc.Search.Streams("Star", opt)

	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}

}

func TestSearchGames(t *testing.T) {

	tc := NewClient(&http.Client{})
	tru := true
	opt := &ListOptions{
		Live: &tru,
	}

	_, err := tc.Search.Games("Diablo", opt)

	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}

}
