package main

import (
	"fmt"

	"github.com/xopenapi/uim-api-go"
)

func main() {
	api := uim.New("YOUR_TOKEN_HERE")

	userID := "USER_ID"

	_, _, channelID, err := api.OpenIMChannel(userID)

	if err != nil {
		fmt.Printf("%s\n", err)
	}

	api.PostMessage(channelID, uim.MsgOptionText("Hello World!", false))
}
