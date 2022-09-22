package discordhooks

// NewWebhook allows you to create a brand new webhooking instance
func NewWebhook(content string) *webhook {
	return &webhook{
		Contents: content,
	}
}

// AppendEmbed will attempt to create a new embed inside the webhook instance
func (W *webhook) AppendEmbeds(embed ...embed) {
	W.Embeds = append(W.Embeds, embed...)
} 

// NewEmbed implements a new webhook instance
func (W *webhook) NewEmbed(title, description, sender string) embed {
	return embed{
		Title:       title,
		Description: description,
		Author:      author{Name: sender},
		
	}
}


// Introduce fields into the embed, You may have multiple fields per embed
func (e *embed) Field(name, value string, inline bool) {
	e.Fields = append(e.Fields, field{Name: name, Value: value, Inline: inline})
}