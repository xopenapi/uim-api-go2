package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/xopenapi/uim-api-go2"
)

func main() {
	attachment := uim.Attachment{
		Color:         "good",
		Fallback:      "You successfully posted by Incoming Webhook URL!",
		AuthorName:    "uim-go/uim",
		AuthorSubname: "github.com",
		AuthorLink:    "https://github.com/xopenapi/uim-api-go2",
		AuthorIcon:    "https://avatars2.githubusercontent.com/u/652790",
		Text:          "<!channel> All text in UIM uses the same system of escaping: chat messages, direct messages, file comments, etc. :smile:\nSee <https://api.uim.com/docs/message-formatting#linking_to_channels_and_users>",
		Footer:        "uim api",
		FooterIcon:    "https://platform.uim-edge.com/img/default_application_icon.png",
		Ts:            json.Number(strconv.FormatInt(time.Now().Unix(), 10)),
	}
	msg := uim.WebhookMessage{
		Attachments: []uim.Attachment{attachment},
	}

	err := uim.PostWebhook("YOUR_WEBHOOK_URL_HERE", &msg)
	if err != nil {
		fmt.Println(err)
	}
}
