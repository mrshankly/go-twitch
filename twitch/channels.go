// Channels methods of the twitch api.
// https://github.com/justintv/Twitch-API/blob/master/v3_resources/channels.md

package twitch

import (
	"fmt"
	"net/url"
	"strconv"
)

// used with GET /channels/:channel/videos
type VideosS struct {
	Videos []VideoS `json:"videos,omitempty"`
	Links  *LinksS  `json:"_links,omitempty"`
}

// used with GET /channels/:channel/editors
type EditorsS struct {
	Users []UserS `json:"users,omitempty"`
	Links *LinksS `json:"_links,omitempty"`
}

// used with GET /channels/:channel/follows
type FollowsS struct {
	Follows []FollowS `json:"follows,omitempty"`
	Total   int       `json:"_total,omitempty"`
	Links   *LinksS   `json:"_links,omitempty"`
}

type FollowS struct {
	User *UserS `json:"user,omitempty"`
}

type SubsS struct {
	Total         int     `json:"_total,omitempty"`
	Links         *LinksS `json:"_links,omitempty"`
	Subscriptions []SubS  `json:"subscriptions,omitempty"`
}

type SubS struct {
	Id   string `json:"_id,omitempty"`
	User *UserS `json:"user,omitempty"`
}

type ChannelsMethod struct {
	client *Client
}

// Returns a channel object. If `name` is an empty string, returns the channel
// object of authenticated user.
func (c *ChannelsMethod) Channel(name string) (*ChannelS, error) {
	rel := "channel" // get authenticated channel
	if name != "" {
		rel = "channels/" + name
	}

	channel := new(ChannelS)
	_, err := c.client.Get(rel, channel)
	return channel, err
}

// Returns a list of users who are editors of channel `name`.
func (c *ChannelsMethod) Editors(name string) (*EditorsS, error) {
	rel := "channels/" + name + "/editors"

	editors := new(EditorsS)
	_, err := c.client.Get(rel, editors)
	return editors, err
}

// Returns a list of videos ordered by time of creation, starting with the most
// recent from channel `name`.
func (c *ChannelsMethod) Videos(name string, opt *ListOptions) (*VideosS, error) {
	rel := "channels/" + name + "/videos"
	if opt != nil {
		p := url.Values{
			"limit":  []string{strconv.Itoa(opt.Limit)},
			"offset": []string{strconv.Itoa(opt.Offset)},
		}
		rel += "?" + p.Encode()
	}

	videos := new(VideosS)
	_, err := c.client.Get(rel, videos)
	return videos, err
}

// Returns a list of users the channel `name` is following.
func (c *ChannelsMethod) Follows(name string, opt *ListOptions) (*FollowsS, error) {
	rel := "channels/" + name + "/follows"
	if opt != nil {
		p := url.Values{
			"limit":     []string{strconv.Itoa(opt.Limit)},
			"offset":    []string{strconv.Itoa(opt.Offset)},
			"direction": []string{opt.Direction},
		}
		rel += "?" + p.Encode()
	}

	follow := new(FollowsS)
	_, err := c.client.Get(rel, follow)
	return follow, err
}

func (c *ChannelsMethod) Subscriptions(name string, opt *ListOptions) (*SubsS, error) {
	rel := "channels/" + name + "/subscriptions"
	if opt != nil {
		p := url.Values{
			"limit":     []string{strconv.Itoa(opt.Limit)},
			"offset":    []string{strconv.Itoa(opt.Offset)},
			"direction": []string{opt.Direction},
		}
		rel += "?" + p.Encode()
	}

	subs := new(SubsS)
	_, err := c.client.Get(rel, subs)
	return subs, err
}

func (c *ChannelsMethod) Subscription(name string, user string) (*SubS, error) {
	rel := fmt.Sprintf("channels/%s/subscriptions/%s", name, user)

	sub := new(SubS)
	_, err := c.client.Get(rel, sub)
	return sub, err
}

// TODO PUT /channels/:channel/

// TODO POST /channels/:channel/commercial
