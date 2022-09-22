package main

import (
	"discordhooks"
)

func main() {
	instance := discordhooks.New(, "")


	/* Makes a webhook */
	model := discordhooks.NewWebhook("")
	
	/* Makes a embed */
	embed := model.NewEmbed("An embed has spawned", "I'm an embed description", "Spawner", )
	embed.Field("1", "my 1st webhook field", false)
	embed.Field("2", "my 2nd webhook field", true)

	embed.Footer = discordhooks.Footer{
		Text: "hello, world!",
	}

	/* Thumbnail for the webhook */
	embed.Thumbnail = discordhooks.Thumbnail{
		URL: "https://res.cloudinary.com/practicaldev/image/fetch/s--AgjcyHDM--/c_imagga_scale,f_auto,fl_progressive,h_420,q_auto,w_1000/https://dev-to-uploads.s3.amazonaws.com/i/8y9m1r90a9moi4ufe6lm.png",
	}

	model.AppendEmbeds(embed)

	/* Attempts to send the webhook request */
	if err := instance.NewWebhook(model); err != nil {
		panic(err)
	}
}
