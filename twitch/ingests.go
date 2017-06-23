package twitch

import (
	"context"
	"net/http"
)

const ingestsPath = "ingests"

// IngestsService handles communication with the ingests related
// methods of the Twitch API.
//
// Twitch API docs: https://dev.twitch.tv/docs/v5/reference/ingests/
type IngestsService service

// Ingest represents an ingest Twitch server.
type Ingest struct {
	ID           int     `json:"_id,omitempty"`
	Availability float64 `json:"availability,omitempty"`
	Default      bool    `json:"default,omitempty"`
	Name         string  `json:"name,omitempty"`
	URLTemplate  string  `json:"url_template,omitempty"`
}

type ingestsRoot struct {
	Ingests []*Ingest `json:"ingests,omitempty"`
}

// Returns a list of Twitch ingest servers.
//
// Twitch API docs: https://dev.twitch.tv/docs/v5/reference/ingests/#get-ingest-server-list
func (s *IngestsService) List(ctx context.Context) ([]*Ingest, *http.Response, error) {
	req, err := s.client.NewRequest("GET", ingestsPath, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(ingestsRoot)
	resp, err := s.client.Do(ctx, req, &root)
	if err != nil {
		return nil, resp, err
	}

	return root.Ingests, resp, nil
}
