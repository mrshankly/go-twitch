package twitch

import (
	"context"
	"fmt"
	"net/http"
)

const chatPath = "chat"

// ChatService handles communication with the chat related methods of the
// Twitch API.
//
// Twitch API docs: https://dev.twitch.tv/docs/v5/reference/chat/
type ChatService service

type Badge struct {
	Alpha string `json:"alpha,omitempty"`
	Image string `json:"image,omitempty"`
	Svg   string `json:"svg,omitempty"`
}

type Badges struct {
	Admin       *Badge `json:"admin,omitempty"`
	Broadcaster *Badge `json:"broadcaster,omitempty"`
	GlobalMod   *Badge `json:"global_mod,omitempty"`
	Mod         *Badge `json:"mod,omitempty"`
	Staff       *Badge `json:"staff,omitempty"`
	Subscriber  *Badge `json:"subscriber"`
	Turbo       *Badge `json:"turbo,omitempty"`
}

type EmoticonSimple struct {
	Code        string `json:"code,omitempty"`
	EmoticonSet int    `json:"emoticon_set,omitempty"`
	ID          int    `json:"id,omitempty"`
}

type Emoticon struct {
	Regex  string `json:"regex,omitempty"`
	Images []*struct {
		Width       int    `json:"width,omitempty"`
		Height      int    `json:"height,omitempty"`
		URL         string `json:"url,omitempty"`
		EmoticonSet int    `json:"emoticon_set,omitempty"`
	} `json:"images,omitempty"`
}

type ChatOptions struct {
	EmoteSets []int `url:"emotesets,omitempty,comma"`
}

type emoticonsBySetRoot struct {
	Emoticons map[int][]*EmoticonSimple `json:"emoticon_sets,omitempty"`
}

type emoticonsRoot struct {
	Emoticons []*Emoticon `json:"emoticons,omitempty"`
}

// Returns the badges that can be used in the chat of the specified channel.
//
// Twitch API docs: https://dev.twitch.tv/docs/v5/reference/chat/#get-chat-badges-by-channel
func (s *ChatService) Badges(ctx context.Context, channel string) (*Badges, *http.Response, error) {
	path := fmt.Sprintf("%v/%v/badges", chatPath, channel)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	badges := new(Badges)
	resp, err := s.client.Do(ctx, req, badges)
	if err != nil {
		return nil, resp, err
	}

	return badges, resp, nil
}

// Returns all chat emoticons (not including their images) in one or more
// specified sets.
//
// Twitch API docs: https://dev.twitch.tv/docs/v5/reference/chat/#get-chat-emoticons-by-set
func (s *ChatService) EmoticonsBySet(ctx context.Context, opt *ChatOptions) (map[int][]*EmoticonSimple, *http.Response, error) {
	path := fmt.Sprintf("%v/emoticon_images", chatPath)
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(emoticonsBySetRoot)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.Emoticons, resp, nil
}

// Returns all chat emoticons (including their images).
//
// Twitch API docs: https://dev.twitch.tv/docs/v5/reference/chat/#get-all-chat-emoticons
func (s *ChatService) Emoticons(ctx context.Context) ([]*Emoticon, *http.Response, error) {
	path := fmt.Sprintf("%v/emoticons", chatPath)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(emoticonsRoot)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.Emoticons, resp, nil
}
