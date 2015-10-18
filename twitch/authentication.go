package twitch

import (
	"net/http"
	"time"
	"github.com/google/go-querystring/query"
	"net/url"
	"bytes"
	"strconv"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type AuthMethods struct {
	client *Client
	AccessToken string
}

type UserInfo struct {
	Display_Name  string
	Id            uint `json:"_id"`
	Name          string
	Type          string
	Bio           string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Logo          string
	Email         string
	Partnered     bool
	Notifications struct {
		Push  bool
		Email bool
	}
}

type OAuthResponse struct {
	Access_Token  string
	Refresh_Token string
	Scope         []string
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
	fmt.Println(url)

	http.Redirect(w, r, url, http.StatusFound)
}

func (a *AuthMethods) LoginCallback(w http.ResponseWriter, r *http.Request) (*OAuthResponse,error){

	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Add("client_id", a.client.AppInfo.ClientID)
	data.Add("client_secret", a.client.AppInfo.ClientSecret)
	data.Add("redirect_uri", a.client.AppInfo.RedirectURI)
	data.Add("state", a.client.AppInfo.State)
	fmt.Println(r.URL.Query())
	code := r.URL.Query().Get("code")
	data.Add("code", code)

	req, err := http.NewRequest("POST", rootURL + "oauth2/token", bytes.NewBufferString(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	fmt.Println(rootURL + "oauth2/token")
	fmt.Println(data)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error 1")
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Println(resp)

	if resp.StatusCode != 200 {
		fmt.Println("Error 2")
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error 3")
		return nil, err
	}

	oauth := &OAuthResponse{}

	if err := json.Unmarshal(body, oauth); err != nil {
		fmt.Println("Error 4")
		return nil, err
	}

	fmt.Println(oauth)

	return oauth, nil

}