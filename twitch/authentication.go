package twitch

import (
	"bytes"
	"encoding/json"
	"github.com/google/go-querystring/query"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type AuthMethods struct {
	client      *Client
	AccessToken string
}

type OAuthResponse struct {
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
	Scope        []string `json:"scope"`
}

type OAuthCallback func(http.ResponseWriter, *http.Request) (*OAuthResponse, error)

func (a *AuthMethods) OAuthRedirect(w http.ResponseWriter, r *http.Request) {

	url := rootURL + "oauth2/authorize"

	params := *a.client.AppInfo
	params.ClientSecret = ""

	v, err := query.Values(params)
	if err != nil {
		http.Error(w, "Error occured at OAuth Initialiazation", 400)
	}

	url += "?response_type=code&" + v.Encode()

	http.Redirect(w, r, url, http.StatusFound)
}

func (a *AuthMethods) LoginCallback(w http.ResponseWriter, r *http.Request) (*OAuthResponse, error) {

	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Add("client_id", a.client.AppInfo.ClientID)
	data.Add("client_secret", a.client.AppInfo.ClientSecret)
	data.Add("redirect_uri", a.client.AppInfo.RedirectURI)
	data.Add("state", a.client.AppInfo.State)

	code := r.URL.Query().Get("code")
	data.Add("code", code)

	req, err := http.NewRequest("POST", rootURL+"oauth2/token", bytes.NewBufferString(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	oauth := &OAuthResponse{}

	if err := json.Unmarshal(body, oauth); err != nil {
		return nil, err
	}

	return oauth, nil

}
