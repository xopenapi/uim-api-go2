package main

import (
	"flag"
	"fmt"

	"github.com/xopenapi/uim-api-go2"
)

func main() {
	var (
		apiToken string
		debug    bool
	)

	flag.StringVar(&apiToken, "token", "YOUR_TOKEN_HERE", "Your UIM API Token")
	flag.BoolVar(&debug, "debug", false, "Show JSON output")
	flag.Parse()

	api := uim.New(apiToken, uim.OptionDebug(debug))

	// Get all stars for the usr.
	params := uim.NewStarsParameters()
	starredItems, _, err := api.GetStarred(params)
	if err != nil {
		fmt.Printf("Error getting stars: %s\n", err)
		return
	}
	for _, s := range starredItems {
		var desc string
		switch s.Type {
		case uim.TYPE_MESSAGE:
			desc = s.Message.Text
		case uim.TYPE_FILE:
			desc = s.File.Name
		case uim.TYPE_FILE_COMMENT:
			desc = s.File.Name + " - " + s.Comment.Comment
		case uim.TYPE_CHANNEL, uim.TYPE_IM, uim.TYPE_GROUP:
			desc = s.Channel
		}
		fmt.Printf("Starred %s: %s\n", s.Type, desc)
	}
}
