package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/xopenapi/uim-api-go"
)

func main() {
	api := uim.New(
		"YOUR TOKEN HERE",
		uim.OptionDebug(true),
		uim.OptionLog(log.New(os.Stdout, "uim-bot: ", log.Lshortfile|log.LstdFlags)),
	)

	// turn on the batch_presence_aware option
	rtm := api.NewRTM(uim.RTMOptionConnParams(url.Values{
		"batch_presence_aware": {"1"},
	}))
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		fmt.Print("Event Received: ")
		switch ev := msg.Data.(type) {
		case *uim.HelloEvent:
			// Replace USER-ID-N here with your User IDs
			rtm.SendMessage(rtm.NewSubscribeUserPresence([]string{
				"USER-ID-1",
				"USER-ID-2",
			}))

		case *uim.ConnectedEvent:
			fmt.Println("Infos:", ev.Info)
			fmt.Println("Connection counter:", ev.ConnectionCount)
			// Replace C2147483705 with your Channel ID
			rtm.SendMessage(rtm.NewOutgoingMessage("Hello world", "C2147483705"))

		case *uim.MessageEvent:
			fmt.Printf("Message: %v\n", ev)

		case *uim.PresenceChangeEvent:
			fmt.Printf("Presence Change: %v\n", ev)

		case *uim.LatencyReport:
			fmt.Printf("Current latency: %v\n", ev.Value)

		case *uim.RTMError:
			fmt.Printf("Error: %s\n", ev.Error())

		case *uim.InvalidAuthEvent:
			fmt.Printf("Invalid credentials")
			return

		default:

			// Ignore other events..
			// fmt.Printf("Unexpected: %v\n", msg.Data)
		}
	}
}
