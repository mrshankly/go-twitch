// Streams methods of the twitch api.
// https://github.com/justintv/Twitch-API/blob/master/v3_resources/chat.md

package twitch

type ChatLinks struct {
	Links ChatLinksS `json:"_links,omitempty"`
}

type ChatLinksS struct {
	Emoticons string `json:"emoticons,omitempty"`
	Badges    string `json:"badges,omitempty"`
}

type EmoticonsS struct {
	Emoticons []EmoticonS `json:"emoticons,omitempty"`
}

type EmoticonS struct {
	Regex  string   `json:"regex,omitempty"`
	Images []ImageS `json:"images,omitempty"`
}

type ImageS struct {
	EmoticonSet int    `json:"emoticon_set,omitempty"`
	Height      int    `json:"height,omitempty"`
	Width       int    `json:"width,omitempty"`
	Url         string `json:"url,omitempty"`
}

type ChatMethod struct {
	client *Client
}

func (c *ChatMethod) Channel(name string) (*ChatLinks, error) {
	rel := "chat/" + name

	chatLinks := new(ChatLinks)
	_, err := c.client.Get(rel, chatLinks)
	return chatLinks, err
}

func (c *ChatMethod) Emoticons() (*EmoticonsS, error) {
	rel := "chat/emoticons"

	emoticons := new(EmoticonsS)
	_, err := c.client.Get(rel, emoticons)
	return emoticons, err
}
