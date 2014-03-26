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
		p := url.Values{
			"limit":  []string{strconv.Itoa(opt.Limit)},
			"offset": []string{strconv.Itoa(opt.Offset)},
			"game":   []string{opt.Game},
			"period": []string{opt.Period},
		}
		rel += "?" + p.Encode()
	}

	videos := new(VideosS)
	_, err := v.client.Get(rel, videos)
	return videos, err
}
