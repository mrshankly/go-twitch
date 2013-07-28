// Streams methods of the twitch api.
// https://github.com/justintv/Twitch-API/blob/master/v3_resources/chat.md

package twitch

type ChatLinks struct {
	Emoticons string `json:"emoticons,omitempty"`
	Badges    string `json:"badges,omitempty"`
}

type EmoticonS struct {
	Emoticons []*EmoteS `json:"emoticons,omitempty"`
}

type EmoteS struct {
	Regex string `json:"regex,omitempty"`
	Images []*ImageS `json:"images,omitempty"`
}

type ImageS struct {
	EmoticonSet int    `json:"emoticon_set,omitempty"`
	Height      int    `json:"height,omitempty"`
	Width       int    `json:"width,omitempty"`
	Url         string `json:"url,omitempty"`
}

func (c *Method) Channel(name string) (*ChatLinks, error) {
	rel := "chat/" + name

	chatLinks := new(ChatLinks)
	_, err := c.client.Get(rel, chatLinks)
	return chatLinks, err
}

func (c *Method) Emoticons() (*EmoticonsS, error) {
	rel := "chat/emoticons"

	emoticons := new(EmoticonsS)
	_, err := c.client.Get(rel, emoticons)
	return emoticons, err
}
