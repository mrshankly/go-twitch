package twitch

import (
	"context"
	"fmt"
	"net/http"
)

const bitsPath = "bits"

// BitsService handles communication with the bits related methods of the
// Twitch API.
//
// Twitch API docs: https://dev.twitch.tv/docs/v5/reference/bits/
type BitsService service

type CheermoteScales map[string]string

type CheermoteStates map[string]CheermoteScales

type CheermoteImages map[string]CheermoteStates

type Cheermote struct {
	Prefix      string   `json:"prefix,omitempty"`
	Backgrounds []string `json:"backgrounds,omitempty"`
	Scales      []string `json:"scales,omitempty"`
	States      []string `json:"states,omitempty"`
	Tiers       []*struct {
		Color   string          `json:"color,omitempty"`
		ID      string          `json:"id,omitempty"`
		Images  CheermoteImages `json:"images,omitempty"`
		MinBits int             `json:"min_bits,omitempty"`
	} `json:"tiers,omitempty"`
}

type CheermotesListOptions struct {
	ChannelID string `url:"channel_id,omitempty"`
}

type cheermotesRoot struct {
	Actions []*Cheermote `json:"actions,omitempty"`
}

// Returns the list of available cheermotes, animated emotes to which viewers
// can assign bits, to cheer in chat.
//
// Twitch API docs: https://dev.twitch.tv/docs/v5/reference/bits/#get-cheermotes
func (s *BitsService) Cheermotes(ctx context.Context, opt *CheermotesListOptions) ([]*Cheermote, *http.Response, error) {
	path := fmt.Sprintf("%v/actions", bitsPath)
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(cheermotesRoot)
	resp, err := s.client.Do(ctx, req, &root)
	if err != nil {
		return nil, resp, err
	}

	return root.Actions, resp, nil
}
