package twitch

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthRedirect(t *testing.T) {

	tc := NewClient(&http.Client{})

	tc.AppInfo.ClientID = "12"
	tc.AppInfo.ClientSecret = "34"
	tc.AppInfo.State = "56"
	tc.AppInfo.RedirectURI = "http://example.org/users/login_callback"
	tc.AppInfo.Scope = "user_read channel_read"

	expectedRedirect := "https://api.twitch.tv/kraken/oauth2/authorize?response_type=code&client_id=12&client_secret=&" +
		"redirect_uri=http%3A%2F%2Fexample.org%2Fusers%2Flogin_callback&scope=user_read+channel_read&state=56"

	ts := httptest.NewServer(http.HandlerFunc(tc.Auth.OAuthRedirect))
	defer ts.Close()

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if req.URL.String() != expectedRedirect {
				t.FailNow()
				return nil
			} else {
				return errors.New("OK")
			}
		},
	}

	client.Get(ts.URL)
}
