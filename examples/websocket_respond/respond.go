package main

import (
	"fmt"
	"strings"

	"github.com/xopenapi/uim-api-go2"
)

func main() {
	api := uim.New(
		"YOUR-TOKEN-HERE",
	)

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *uim.MessageEvent:
			msg := ev.Msg

			if msg.SubType != "" {
				break // We're only handling normal messages.
			}

			// Create a response object.
			resp := rtm.NewOutgoingMessage(fmt.Sprintf("echo %s", msg.Text), msg.Channel)

			// Respond in thread if not a direct message.
			if !strings.HasPrefix(msg.Channel, "D") {
				resp.ThreadTimestamp = msg.Timestamp
			}

			// Respond in same thread if message came from a thread.
			if msg.ThreadTimestamp != "" {
				resp.ThreadTimestamp = msg.ThreadTimestamp
			}

			rtm.SendMessage(resp)

		case *uim.ConnectedEvent:
			fmt.Println("Connected to UIM")

		case *uim.InvalidAuthEvent:
			fmt.Println("Invalid token")
			return
		}
	}
}
