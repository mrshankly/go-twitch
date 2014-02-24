package twitch

type IngestsS struct {
	Ingests []IngestS `json:"ingests,omitempty"`
}

type IngestS struct {
	Name         string  `json:"name,omitempty"`
	Default      bool    `json:"default,omitempty"`
	Id           int     `json:"_id,omitempty"`
	UrlTemplate  string  `json:"url_template,omitempty"`
	Availability float64 `json:"availability,omitempty"`
}

type IngestsMethod struct {
	client *Client
}

func (i *IngestsMethod) List() (*IngestsS, error) {
	rel := "ingests"

	ingests := new(IngestsS)
	_, err := i.client.Get(rel, ingests)
	return ingests, err
}
