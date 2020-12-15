package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/xopenapi/uim-api-go2"
)

func main() {
	api := uim.New("YOUR_TOKEN_HERE")
	attachment := uim.Attachment{
		Pretext:    "pretext",
		Fallback:   "We don't currently support your client",
		CallbackID: "accept_or_reject",
		Color:      "#3AA3E3",
		Actions: []uim.AttachmentAction{
			uim.AttachmentAction{
				Name:  "accept",
				Text:  "Accept",
				Type:  "button",
				Value: "accept",
			},
			uim.AttachmentAction{
				Name:  "reject",
				Text:  "Reject",
				Type:  "button",
				Value: "reject",
				Style: "danger",
			},
		},
	}

	message := uim.MsgOptionAttachments(attachment)
	channelID, timestamp, err := api.PostMessage("CHANNEL_ID", uim.MsgOptionText("", false), message)
	if err != nil {
		fmt.Printf("Could not send message: %v", err)
	}
	fmt.Printf("Message with buttons sucessfully sent to channel %s at %s", channelID, timestamp)
	http.HandleFunc("/actions", actionHandler)
}

func actionHandler(w http.ResponseWriter, r *http.Request) {
	var payload uim.InteractionCallback
	err := json.Unmarshal([]byte(r.FormValue("payload")), &payload)
	if err != nil {
		fmt.Printf("Could not parse action response JSON: %v", err)
	}
	fmt.Printf("Message button pressed by user %s with value %s", payload.User.Name, payload.Value)
}
