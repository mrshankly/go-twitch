package twitch

import (
	"net/http"
	"testing"
)

func TestChatChannel(t *testing.T) {

	tc := NewClient(&http.Client{})

	_, err := tc.Chat.Channel("kraken_test_user")

	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}

}

func TestChatEmoticons(t *testing.T) {

	tc := NewClient(&http.Client{})

	_, err := tc.Chat.Emoticons()

	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}

}
