package main

import (
	"fmt"

	"github.com/xopenapi/uim-api-go"
)

func main() {
	api := uim.New("YOUR_TOKEN_HERE")
	attachment := uim.Attachment{
		Pretext: "some pretext",
		Text:    "some text",
		// Uncomment the following part to send a field too
		/*
			Fields: []uim.AttachmentField{
				uim.AttachmentField{
					Title: "a",
					Value: "no",
				},
			},
		*/
	}

	channelID, timestamp, err := api.PostMessage("CHANNEL_ID", uim.MsgOptionText("Some text", false), uim.MsgOptionAttachments(attachment))
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}
