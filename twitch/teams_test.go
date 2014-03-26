package twitch

import (
	"net/http"
	"testing"
)

func TestTeamsTeam(t *testing.T) {
	tc := NewClient(&http.Client{})

	_, err := tc.Teams.Team("testteam")

	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}

}

func TestTeamsList(t *testing.T) {
	tc := NewClient(&http.Client{})

	opt := &ListOptions{
		Limit:  1,
		Offset: 0,
	}

	_, err := tc.Teams.List(opt)

	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}

}
