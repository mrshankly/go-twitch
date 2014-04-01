package twitch

import (
	"net/url"
	"strconv"
)

type VideosMethod struct {
	client *Client
}

func (v *VideosMethod) Id(id string) (*VideoS, error) {
	rel := "videos/" + id

	video := new(VideoS)
	_, err := v.client.Get(rel, video)
	return video, err
}

func (v *VideosMethod) Top(opt *ListOptions) (*VideosS, error) {
	rel := "videos/top"
	if opt != nil {
		p := url.Values{}
		if len(opt.Game) > 0 {
			p.Add("game", opt.Game)
		}
		if len(opt.Period) > 0 {
			p.Add("period", opt.Period)
		}
		if opt.Limit > 0 {
			p.Add("limit", strconv.Itoa(opt.Limit))
		}
		if opt.Offset > 0 {
			p.Add("offset", strconv.Itoa(opt.Offset))
		}
		if len(p) > 0 {
			rel += "?" + p.Encode()
		}
	}

	videos := new(VideosS)
	_, err := v.client.Get(rel, videos)
	return videos, err
}
