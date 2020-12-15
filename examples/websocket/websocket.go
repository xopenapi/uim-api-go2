package main

import (
	"fmt"
	"log"
	"os"

	"github.com/xopenapi/uim-api-go2"
)

func main() {
	api := uim.New(
		"YOUR TOKEN HERE",
		uim.OptionDebug(true),
		uim.OptionLog(log.New(os.Stdout, "uim-bot: ", log.Lshortfile|log.LstdFlags)),
	)

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		fmt.Print("Event Received: ")
		switch ev := msg.Data.(type) {
		case *uim.HelloEvent:
			// Ignore hello

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

		case *uim.DesktopNotificationEvent:
			fmt.Printf("Desktop Notification: %v\n", ev)

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
