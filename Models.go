package discordhooks

// Webhook contains information about the current webhook sender and receiver
type webhook struct {
	Contents string		`json:"content"`
	Embeds   []embed    `json:"embeds"`
}

// Formula for the embed fields on webhooks
type embed struct {
	Title       string    `json:"title"`
	Description string	  `json:"description"`
	Colour      int       `json:"color"`
	Fields      []field   `json:"fields"`
	Author      author    `json:"author"`
	Footer      Footer    `json:"footer"`
	Thumbnail   Thumbnail `json:"thumbnail"`
}

// Formula for fields inside the embed on webhooks
type field struct {
	Name string `json:"name"`
	Value string `json:"value"`
	Inline bool `json:"inline"`
}

// author fills the field inside the embed on webhooks for the author who sent the message
type author struct {
	Name string `json:"name"`
}

// footer is placed at the bottom of the embedded webhook request
type Footer struct {
	Text    string `json:"text"`
	IconURL string `json:"icon_url"`
}

// Thumbnail is comprised of only a URL for the webhook
type Thumbnail struct {
	URL string `json:"url"`
}