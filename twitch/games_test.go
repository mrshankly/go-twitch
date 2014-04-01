package twitch

import (
	"net/http"
	"testing"
)

func TestGamesTop(t *testing.T) {
	tc := NewClient(&http.Client{})

	opt := &ListOptions{
		Limit:  1,
		Offset: 0,
		Hls:    &Hls{},
	}

	_, err := tc.Games.Top(opt)

	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}

}
