package twitch

import "github.com/google/go-querystring/query"

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
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "?" + v.Encode()
	}

	videos := new(VideosS)
	_, err := v.client.Get(rel, videos)
	return videos, err
}
