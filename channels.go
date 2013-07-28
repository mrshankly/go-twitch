// Channels methods of the twitch api.
// https://github.com/justintv/Twitch-API/blob/master/v3_resources/streams.md

package twitch

import (
	"strconv"
	"net/url"
)

type VideosS struct {
	Videos []*VideoS `json:"videos,omitempty"`
	Links  *LinksS   `json:"_links,omitempty"`
}

type ChannelsMethod struct {
	client *Client
}

func (c *ChannelsMethod) Channel(name string) (*ChannelS, error) {
	rel := "channels/" + name

	channel := new(ChannelS)
	_, err := c.client.Get(rel, channel)
	return channel, err
}

// TODO GET /channel

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

// TODO PUT /channels/:channel/

// TODO POST /channels/:channel/commercial
