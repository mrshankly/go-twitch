package twitch

import (
	"net/http"
	"testing"
)

func TestIngestsList(t *testing.T) {
	tc := NewClient(&http.Client{})

	_, err := tc.Ingests.List()

	if err != nil {
		t.Errorf("error not nil: ", err)
	}

}
