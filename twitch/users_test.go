package twitch

import (
	"net/http"
	"testing"
)

func TestUsersUser(t *testing.T) {

	tc := NewClient(&http.Client{})

	_, err := tc.Users.User("test_user1")

	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}

}

func TestUsersFollows(t *testing.T) {

	tc := NewClient(&http.Client{})

	opt := &ListOptions{
		Limit:  1,
		Offset: 0,
	}

	_, err := tc.Users.Follows("test_user1", opt)

	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}

}

func TestUsersFollow(t *testing.T) {

	tc := NewClient(&http.Client{})

	_, err := tc.Users.Follow("Roybot1911", "Dansgaming")

	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}

}
