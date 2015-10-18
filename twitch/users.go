package twitch

import (
	"fmt"
	"github.com/google/go-querystring/query"
)

type UAuthedS struct {
	UserS
	Email         string         `json:"email,omitempty"`
	Partnered     bool           `json:"partnered,omitempty"`
	Notifications UNotificationS `json:"notifications,omitempty"`
}

type BlocksS struct {
	Blocks []BlockS `json:"blocks,omitempty"`
	Links  LinksS   `json:"_links,omitempty"`
}

type BlockS struct {
	User UserS `json:"user,omitempty"`
	Id   int   `json:"_id,omitempty"`
}

type UFollowsS struct {
	Follows []UFollowS `json:"follows,omitempty"`
	Links   LinksS     `json:"_links,omitempty"`
	Total   int        `json:"_total,omitempty"`
}

type UFollowS struct {
	Channel ChannelS `json:"channel,omitempty"`
}

type UTargetS struct {
	Channel   ChannelS `json:"channel,omitempty"`
	CreatedAt string   `json:"created_at,omitempty"`
}

type UNotificationS struct {
	Push  bool `json:"push,omitempty"`
	Email bool `json:"email,omitempty"`
}

type UsersMethod struct {
	client *Client
}

// User returns a user object.
func (u *UsersMethod) User(user string) (*UserS, error) {
	rel := "users/" + user

	usr := new(UserS)
	_, err := u.client.Get(rel, usr)
	return usr, err
}

// User returns a user object.
func (u *UsersMethod) Authenticated() (*UAuthedS, error) {
	rel := "user"

	usr := new(UAuthedS)
	_, err := u.client.Get(rel, usr)
	return usr, err
}

func (u *UsersMethod) blocks(login string, opt *ListOptions) (*BlocksS, error) {
	rel := "users/" + login + "/blocks"
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "?" + v.Encode()
	}

	blocks := new(BlocksS)
	_, err := u.client.Get(rel, blocks)
	return blocks, err
}

// Get a user's list of followed channels
func (u *UsersMethod) Follows(user string, opt *ListOptions) (*UFollowsS, error) {
	rel := "users/" + user + "/follows/channels"
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "?" + v.Encode()
	}

	follows := new(UFollowsS)
	_, err := u.client.Get(rel, follows)
	return follows, err
}

// Get status of follow relationship between user and target channel
func (u *UsersMethod) Follow(user, target string) (*UTargetS, error) {
	rel := fmt.Sprintf("users/%s/follows/channels/%s", user, target)

	follow := new(UTargetS)
	_, err := u.client.Get(rel, follow)
	return follow, err
}

func (u *UsersMethod) subscription(user, channel string) (*UTargetS, error) {
	rel := fmt.Sprintf("users/%s/subscriptions/%s", user, channel)

	follow := new(UTargetS)
	_, err := u.client.Get(rel, follow)
	return follow, err
}
