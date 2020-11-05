package main

import (
	"fmt"

	"github.com/xopenapi/uim-api-go"
)

// A minimal example showing how to open a modal
func main() {

	// Create a ModalViewRequest with a header and two inputs
	titleText := uim.NewTextBlockObject("plain_text", "My App", false, false)
	closeText := uim.NewTextBlockObject("plain_text", "Close", false, false)
	submitText := uim.NewTextBlockObject("plain_text", "Submit", false, false)

	headerText := uim.NewTextBlockObject("mrkdwn", "Please enter your name", false, false)
	headerSection := uim.NewSectionBlock(headerText, nil, nil)

	firstNameText := uim.NewTextBlockObject("plain_text", "First Name", false, false)
	firstNamePlaceholder := uim.NewTextBlockObject("plain_text", "Enter your first name", false, false)
	firstNameElement := uim.NewPlainTextInputBlockElement(firstNamePlaceholder, "firstName")
	// Notice that blockID is a unique identifier for a block
	firstName := uim.NewInputBlock("First Name", firstNameText, firstNameElement)

	lastNameText := uim.NewTextBlockObject("plain_text", "Last Name", false, false)
	lastNamePlaceholder := uim.NewTextBlockObject("plain_text", "Enter your first name", false, false)
	lastNameElement := uim.NewPlainTextInputBlockElement(lastNamePlaceholder, "firstName")
	lastName := uim.NewInputBlock("Last Name", lastNameText, lastNameElement)

	blocks := uim.Blocks{
		BlockSet: []uim.Block{
			headerSection,
			firstName,
			lastName,
		},
	}

	var modalRequest uim.ModalViewRequest
	modalRequest.Type = uim.ViewType("modal")
	modalRequest.Title = titleText
	modalRequest.Close = closeText
	modalRequest.Submit = submitText
	modalRequest.Blocks = blocks

	api := uim.New("YOUR_BOT_TOKEN_HERE")

	// Using a trigger ID you can open a modal
	// The trigger ID is provided through certain events and interactions
	// More information can be found here: https://api.uim.com/interactivity/handling#modal_responses
	_, err := api.OpenView("YOUR_TRIGGERID_HERE", modalRequest)
	if err != nil {
		fmt.Printf("Error opening view: %s", err)
	}
}
