package twitch

import (
	"net/url"
	"strconv"
)

type VideosMethod struct {
	client *Client
}

func (v *VideosMethod) Id(id int) (*VideoS, error) {
	rel := "videos/" + strconv.Itoa(id)

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
			"games":  []string{opt.Game},
			"period": []string{opt.Period},
		}
		rel += "?" + p.Encode()
	}

	videos := new(VideosS)
	_, err := v.client.Get(rel, videos)
	return videos, err
}

func (v *VideosMethod) Followed(opt *ListOptions) (*VideosS, error) {
	rel := "videos/followed"
	if opt != nil {
		p := url.Values{
			"limit":  []string{strconv.Itoa(opt.Limit)},
			"offset": []string{strconv.Itoa(opt.Offset)},
		}
		rel += "?" + p.Encode()
	}

	follows := new(VideosS)
	_, err := v.client.Get(rel, follows)
	return follows, err
}
