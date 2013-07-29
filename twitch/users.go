package twitch

import (
	"fmt"
	"net/url"
	"strconv"
)

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
}

type UFollowS struct {
	Channel ChannelS `json:"channel,omitempty"`
}

type UTargetS struct {
	Channel ChannelS `json:"channel,omitempty"`
}

type UsersMethod struct {
	client Client
}

func (u *UsersMethod) Channel(user string) (*UserS, error) {
	rel := "user" // get authenticated user
	if user != "" {
		rel = "users/" + user
	}

	usr := new(UserS)
	_, err := u.client.Get(rel, usr)
	return usr, err
}

func (u *UsersMethod) Blocks(login string, opt *ListOptions) (*BlocksS, error) {
	rel := "users/" + login + "/blocks"
	if opt != nil {
		p := url.Values{
			"limit":  []string{strconv.Itoa(opt.Limit)},
			"offset": []string{strconv.Itoa(opt.Offset)},
		}
		rel += "?" + p.Encode()
	}

	blocks := new(BlocksS)
	_, err := u.client.Get(rel, blocks)
	return blocks, err
}

func (u *UsersMethod) Follows(user string, opt *ListOptions) (*UFollowsS, error) {
	rel := "users/" + user + "/follows/channels"
	if opt != nil {
		p := url.Values{
			"limit":     []string{strconv.Itoa(opt.Limit)},
			"offset":    []string{strconv.Itoa(opt.Offset)},
			"direction": []string{opt.Direction},
		}
		rel += "?" + p.Encode()
	}

	follows := new(UFollowsS)
	_, err := u.client.Get(rel, follows)
	return follows, err
}

func (u *UsersMethod) Follow(user, target string) (*UTargetS, error) {
	rel := fmt.Sprintf("users/%s/follows/channels/%s", user, target)

	follow := new(UTargetS)
	_, err := u.client.Get(rel, follow)
	return follow, err
}

func (u *UsersMethod) Subscription(user, channel string) (*UTargetS, error) {
	rel := fmt.Sprintf("users/%s/subscriptions/%s", user, channel)

	follow := new(UTargetS)
	_, err := u.client.Get(rel, follow)
	return follow, err
}
