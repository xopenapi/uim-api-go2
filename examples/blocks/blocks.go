package main

import (
	"encoding/json"
	"fmt"

	"github.com/xopenapi/uim-api-go2"
)

// The functions below mock the different templates uim has as examples on their website.
//
// Refer to README.md for more information on the examples and how to use them.

func main() {

	fmt.Println("--- Begin Example One ---")
	exampleOne()
	fmt.Println("--- End Example One ---")

	fmt.Println("--- Begin Example Two ---")
	exampleTwo()
	fmt.Println("--- End Example Two ---")

	fmt.Println("--- Begin Example Three ---")
	exampleThree()
	fmt.Println("--- End Example Three ---")

	fmt.Println("--- Begin Example Four ---")
	exampleFour()
	fmt.Println("--- End Example Four ---")

	fmt.Println("--- Begin Example Five ---")
	exampleFive()
	fmt.Println("--- End Example Five ---")

	fmt.Println("--- Begin Example Six ---")
	exampleSix()
	fmt.Println("--- End Example Six ---")

	fmt.Println("--- Begin Example Unmarshalling ---")
	unmarshalExample()
	fmt.Println("--- End Example Unmarshalling ---")
}

// approvalRequest mocks the simple "Approval" template located on block kit builder website
func exampleOne() {

	// Header Section
	headerText := uim.NewTextBlockObject("mrkdwn", "You have a new request:\n*<fakeLink.toEmployeeProfile.com|Fred Enriquez - New device request>*", false, false)
	headerSection := uim.NewSectionBlock(headerText, nil, nil)

	// Fields
	typeField := uim.NewTextBlockObject("mrkdwn", "*Type:*\nComputer (laptop)", false, false)
	whenField := uim.NewTextBlockObject("mrkdwn", "*When:*\nSubmitted Aut 10", false, false)
	lastUpdateField := uim.NewTextBlockObject("mrkdwn", "*Last Update:*\nMar 10, 2015 (3 years, 5 months)", false, false)
	reasonField := uim.NewTextBlockObject("mrkdwn", "*Reason:*\nAll vowel keys aren't working.", false, false)
	specsField := uim.NewTextBlockObject("mrkdwn", "*Specs:*\n\"Cheetah Pro 15\" - Fast, really fast\"", false, false)

	fieldSlice := make([]*uim.TextBlockObject, 0)
	fieldSlice = append(fieldSlice, typeField)
	fieldSlice = append(fieldSlice, whenField)
	fieldSlice = append(fieldSlice, lastUpdateField)
	fieldSlice = append(fieldSlice, reasonField)
	fieldSlice = append(fieldSlice, specsField)

	fieldsSection := uim.NewSectionBlock(nil, fieldSlice, nil)

	// Approve and Deny Buttons
	approveBtnTxt := uim.NewTextBlockObject("plain_text", "Approve", false, false)
	approveBtn := uim.NewButtonBlockElement("", "click_me_123", approveBtnTxt)

	denyBtnTxt := uim.NewTextBlockObject("plain_text", "Deny", false, false)
	denyBtn := uim.NewButtonBlockElement("", "click_me_123", denyBtnTxt)

	actionBlock := uim.NewActionBlock("", approveBtn, denyBtn)

	// Build Message with blocks created above

	msg := uim.NewBlockMessage(
		headerSection,
		fieldsSection,
		actionBlock,
	)

	b, err := json.MarshalIndent(msg, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(b))

}

// exampleTwo mocks the more complex "Approval" template located on block kit builder website
// which includes an accessory image next to the approval request
func exampleTwo() {

	// Header Section
	headerText := uim.NewTextBlockObject("mrkdwn", "You have a new request:\n*<google.com|Fred Enriquez - Time Off request>*", false, false)
	headerSection := uim.NewSectionBlock(headerText, nil, nil)

	approvalText := uim.NewTextBlockObject("mrkdwn", "*Type:*\nPaid time off\n*When:*\nAug 10-Aug 13\n*Hours:* 16.0 (2 days)\n*Remaining balance:* 32.0 hours (4 days)\n*Comments:* \"Family in town, going camping!\"", false, false)
	approvalImage := uim.NewImageBlockElement("https://api.uim.com/img/blocks/bkb_template_images/approvalsNewDevice.png", "computer thumbnail")

	fieldsSection := uim.NewSectionBlock(approvalText, nil, uim.NewAccessory(approvalImage))

	// Approve and Deny Buttons
	approveBtnTxt := uim.NewTextBlockObject("plain_text", "Approve", false, false)
	approveBtn := uim.NewButtonBlockElement("", "click_me_123", approveBtnTxt)

	denyBtnTxt := uim.NewTextBlockObject("plain_text", "Deny", false, false)
	denyBtn := uim.NewButtonBlockElement("", "click_me_123", denyBtnTxt)

	actionBlock := uim.NewActionBlock("", approveBtn, denyBtn)

	// Build Message with blocks created above

	msg := uim.NewBlockMessage(
		headerSection,
		fieldsSection,
		actionBlock,
	)

	b, err := json.MarshalIndent(msg, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(b))

}

// exampleThree generates the notification example from the block kit builder website
func exampleThree() {

	// Shared Assets for example
	chooseBtnText := uim.NewTextBlockObject("plain_text", "Choose", true, false)
	chooseBtnEle := uim.NewButtonBlockElement("", "click_me_123", chooseBtnText)
	divSection := uim.NewDividerBlock()

	// Header Section
	headerText := uim.NewTextBlockObject("plain_text", "Looks like you have a scheduling conflict with this event:", false, false)
	headerSection := uim.NewSectionBlock(headerText, nil, nil)

	// Schedule Info Section
	scheduleText := uim.NewTextBlockObject("mrkdwn", "*<fakeLink.toUserProfiles.com|Iris / Zelda 1-1>*\nTuesday, January 21 4:00-4:30pm\nBuilding 2 - Havarti Cheese (3)\n2 guests", false, false)
	scheduleAccessory := uim.NewImageBlockElement("https://api.uim.com/img/blocks/bkb_template_images/notifications.png", "calendar thumbnail")
	schedeuleSection := uim.NewSectionBlock(scheduleText, nil, uim.NewAccessory(scheduleAccessory))

	// Conflict Section
	conflictImage := uim.NewImageBlockElement("https://api.uim.com/img/blocks/bkb_template_images/notificationsWarningIcon.png", "notifications warning icon")
	conflictText := uim.NewTextBlockObject("mrkdwn", "*Conflicts with Team Huddle: 4:15-4:30pm*", false, false)

	conflictSection := uim.NewContextBlock(
		"",
		[]uim.MixedElement{conflictImage, conflictText}...,
	)

	// Proposese Text
	proposeText := uim.NewTextBlockObject("mrkdwn", "*Propose a new time:*", false, false)
	proposeSection := uim.NewSectionBlock(proposeText, nil, nil)

	// Option 1
	optionOneText := uim.NewTextBlockObject("mrkdwn", "*Today - 4:30-5pm*\nEveryone is available: @iris, @zelda", false, false)
	optionOneSection := uim.NewSectionBlock(optionOneText, nil, uim.NewAccessory(chooseBtnEle))

	// Option 2
	optionTwoText := uim.NewTextBlockObject("mrkdwn", "*Tomorrow - 4-4:30pm*\nEveryone is available: @iris, @zelda", false, false)
	optionTwoSection := uim.NewSectionBlock(optionTwoText, nil, uim.NewAccessory(chooseBtnEle))

	// Option 3
	optionThreeText := uim.NewTextBlockObject("mrkdwn", "*Tomorrow - 6-6:30pm*\nSome people aren't available: @iris, ~@zelda~", false, false)
	optionThreeSection := uim.NewSectionBlock(optionThreeText, nil, uim.NewAccessory(chooseBtnEle))

	// Show More Times Link
	showMoreText := uim.NewTextBlockObject("mrkdwn", "*<fakelink.ToMoreTimes.com|Show more times>*", false, false)
	showMoreSection := uim.NewSectionBlock(showMoreText, nil, nil)

	// Build Message with blocks created above
	msg := uim.NewBlockMessage(
		headerSection,
		divSection,
		schedeuleSection,
		conflictSection,
		divSection,
		proposeSection,
		optionOneSection,
		optionTwoSection,
		optionThreeSection,
		showMoreSection,
	)

	b, err := json.MarshalIndent(msg, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(b))

}

// exampleFour profiles a poll example block
func exampleFour() {

	// Shared Assets for example
	divSection := uim.NewDividerBlock()
	voteBtnText := uim.NewTextBlockObject("plain_text", "Vote", true, false)
	voteBtnEle := uim.NewButtonBlockElement("", "click_me_123", voteBtnText)
	profileOne := uim.NewImageBlockElement("https://api.uim.com/img/blocks/bkb_template_images/profile_1.png", "Michael Scott")
	profileTwo := uim.NewImageBlockElement("https://api.uim.com/img/blocks/bkb_template_images/profile_2.png", "Dwight Schrute")
	profileThree := uim.NewImageBlockElement("https://api.uim.com/img/blocks/bkb_template_images/profile_3.png", "Pam Beasely")
	profileFour := uim.NewImageBlockElement("https://api.uim.com/img/blocks/bkb_template_images/profile_4.png", "Angela")

	// Header Section
	headerText := uim.NewTextBlockObject("mrkdwn", "*Where should we order lunch from?* Poll by <fakeLink.toUser.com|Mark>", false, false)
	headerSection := uim.NewSectionBlock(headerText, nil, nil)

	// Option One Info
	optOneText := uim.NewTextBlockObject("mrkdwn", ":sushi: *Ace Wasabi Rock-n-Roll Sushi Bar*\nThe best landlocked sushi restaurant.", false, false)
	optOneSection := uim.NewSectionBlock(optOneText, nil, uim.NewAccessory(voteBtnEle))

	// Option One Votes
	optOneVoteText := uim.NewTextBlockObject("plain_text", "3 votes", true, false)
	optOneContext := uim.NewContextBlock("", []uim.MixedElement{profileOne, profileTwo, profileThree, optOneVoteText}...)

	// Option Two Info
	optTwoText := uim.NewTextBlockObject("mrkdwn", ":hamburger: *Super Hungryman Hamburgers*\nOnly for the hungriest of the hungry.", false, false)
	optTwoSection := uim.NewSectionBlock(optTwoText, nil, uim.NewAccessory(voteBtnEle))

	// Option Two Votes
	optTwoVoteText := uim.NewTextBlockObject("plain_text", "2 votes", true, false)
	optTwoContext := uim.NewContextBlock("", []uim.MixedElement{profileFour, profileTwo, optTwoVoteText}...)

	// Option Three Info
	optThreeText := uim.NewTextBlockObject("mrkdwn", ":ramen: *Kagawa-Ya Udon Noodle Shop*\nDo you like to shop for noodles? We have noodles.", false, false)
	optThreeSection := uim.NewSectionBlock(optThreeText, nil, uim.NewAccessory(voteBtnEle))

	// Option Three Votes
	optThreeVoteText := uim.NewTextBlockObject("plain_text", "No votes", true, false)
	optThreeContext := uim.NewContextBlock("", []uim.MixedElement{optThreeVoteText}...)

	// Suggestions Action
	btnTxt := uim.NewTextBlockObject("plain_text", "Add a suggestion", false, false)
	nextBtn := uim.NewButtonBlockElement("", "click_me_123", btnTxt)
	actionBlock := uim.NewActionBlock("", nextBtn)

	// Build Message with blocks created above
	msg := uim.NewBlockMessage(
		headerSection,
		divSection,
		optOneSection,
		optOneContext,
		optTwoSection,
		optTwoContext,
		optThreeSection,
		optThreeContext,
		divSection,
		actionBlock,
	)

	b, err := json.MarshalIndent(msg, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(b))

}

func exampleFive() {

	// Build Header Section Block, includes text and overflow menu

	headerText := uim.NewTextBlockObject("mrkdwn", "We found *205 Hotels* in New Orleans, LA from *12/14 to 12/17*", false, false)

	// Build Text Objects associated with each option
	overflowOptionTextOne := uim.NewTextBlockObject("plain_text", "Option One", false, false)
	overflowOptionTextTwo := uim.NewTextBlockObject("plain_text", "Option Two", false, false)
	overflowOptionTextThree := uim.NewTextBlockObject("plain_text", "Option Three", false, false)

	// Build each option, providing a value for the option
	overflowOptionOne := uim.NewOptionBlockObject("value-0", overflowOptionTextOne)
	overflowOptionTwo := uim.NewOptionBlockObject("value-1", overflowOptionTextTwo)
	overflowOptionThree := uim.NewOptionBlockObject("value-2", overflowOptionTextThree)

	// Build overflow section
	overflow := uim.NewOverflowBlockElement("", overflowOptionOne, overflowOptionTwo, overflowOptionThree)

	// Create the header section
	headerSection := uim.NewSectionBlock(headerText, nil, uim.NewAccessory(overflow))

	// Shared Divider
	divSection := uim.NewDividerBlock()

	// Shared Objects
	locationPinImage := uim.NewImageBlockElement("https://api.uim.com/img/blocks/bkb_template_images/tripAgentLocationMarker.png", "Location Pin Icon")

	// First Hotel Listing
	hotelOneInfo := uim.NewTextBlockObject("mrkdwn", "*<fakeLink.toHotelPage.com|Windsor Court Hotel>*\n★★★★★\n$340 per night\nRated: 9.4 - Excellent", false, false)
	hotelOneImage := uim.NewImageBlockElement("https://api.uim.com/img/blocks/bkb_template_images/tripAgent_1.png", "Windsor Court Hotel thumbnail")
	hotelOneLoc := uim.NewTextBlockObject("plain_text", "Location: Central Business District", true, false)

	hotelOneSection := uim.NewSectionBlock(hotelOneInfo, nil, uim.NewAccessory(hotelOneImage))
	hotelOneContext := uim.NewContextBlock("", []uim.MixedElement{locationPinImage, hotelOneLoc}...)

	// Second Hotel Listing
	hotelTwoInfo := uim.NewTextBlockObject("mrkdwn", "*<fakeLink.toHotelPage.com|The Ritz-Carlton New Orleans>*\n★★★★★\n$340 per night\nRated: 9.1 - Excellent", false, false)
	hotelTwoImage := uim.NewImageBlockElement("https://api.uim.com/img/blocks/bkb_template_images/tripAgent_2.png", "Ritz-Carlton New Orleans thumbnail")
	hotelTwoLoc := uim.NewTextBlockObject("plain_text", "Location: French Quarter", true, false)

	hotelTwoSection := uim.NewSectionBlock(hotelTwoInfo, nil, uim.NewAccessory(hotelTwoImage))
	hotelTwoContext := uim.NewContextBlock("", []uim.MixedElement{locationPinImage, hotelTwoLoc}...)

	// Third Hotel Listing
	hotelThreeInfo := uim.NewTextBlockObject("mrkdwn", "*<fakeLink.toHotelPage.com|Omni Royal Orleans Hotel>*\n★★★★★\n$419 per night\nRated: 8.8 - Excellent", false, false)
	hotelThreeImage := uim.NewImageBlockElement("https://api.uim.com/img/blocks/bkb_template_images/tripAgent_3.png", "https://api.uim.com/img/blocks/bkb_template_images/tripAgent_3.png")
	hotelThreeLoc := uim.NewTextBlockObject("plain_text", "Location: French Quarter", true, false)

	hotelThreeSection := uim.NewSectionBlock(hotelThreeInfo, nil, uim.NewAccessory(hotelThreeImage))
	hotelThreeContext := uim.NewContextBlock("", []uim.MixedElement{locationPinImage, hotelThreeLoc}...)

	// Action button
	btnTxt := uim.NewTextBlockObject("plain_text", "Next 2 Results", false, false)
	nextBtn := uim.NewButtonBlockElement("", "click_me_123", btnTxt)
	actionBlock := uim.NewActionBlock("", nextBtn)

	// Build Message with blocks created above
	msg := uim.NewBlockMessage(
		headerSection,
		divSection,
		hotelOneSection,
		hotelOneContext,
		divSection,
		hotelTwoSection,
		hotelTwoContext,
		divSection,
		hotelThreeSection,
		hotelThreeContext,
		divSection,
		actionBlock,
	)

	b, err := json.MarshalIndent(msg, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(b))

}

func exampleSix() {

	// Shared Assets for example
	divSection := uim.NewDividerBlock()

	// Shared Available Options
	manageTxt := uim.NewTextBlockObject("plain_text", "Manage", true, false)
	editTxt := uim.NewTextBlockObject("plain_text", "Edit it", false, false)
	readTxt := uim.NewTextBlockObject("plain_text", "Read it", false, false)
	saveTxt := uim.NewTextBlockObject("plain_text", "Save it", false, false)

	editOpt := uim.NewOptionBlockObject("value-0", editTxt)
	readOpt := uim.NewOptionBlockObject("value-1", readTxt)
	saveOpt := uim.NewOptionBlockObject("value-2", saveTxt)

	availableOption := uim.NewOptionsSelectBlockElement("static_select", manageTxt, "", editOpt, readOpt, saveOpt)

	// Header Section
	headerText := uim.NewTextBlockObject("mrkdwn", ":mag: Search results for *Cata*", false, false)
	headerSection := uim.NewSectionBlock(headerText, nil, nil)

	// Result One
	resultOneTxt := uim.NewTextBlockObject("mrkdwn", "*<fakeLink.toYourApp.com|Use Case Catalogue>*\nUse Case Catalogue for the following departments/roles...", false, false)
	resultOneSection := uim.NewSectionBlock(resultOneTxt, nil, uim.NewAccessory(availableOption))

	// Result Two
	resultTwoTxt := uim.NewTextBlockObject("mrkdwn", "*<fakeLink.toYourApp.com|Customer Support - Workflow Diagram Catalogue>*\nThis resource was put together by members of...", false, false)
	resultTwoSection := uim.NewSectionBlock(resultTwoTxt, nil, uim.NewAccessory(availableOption))

	// Result Three
	resultThreeTxt := uim.NewTextBlockObject("mrkdwn", "*<fakeLink.toYourApp.com|Self-Serve Learning Options Catalogue>*\nSee the learning and development options we...", false, false)
	resultThreeSection := uim.NewSectionBlock(resultThreeTxt, nil, uim.NewAccessory(availableOption))

	// Result Four
	resultFourTxt := uim.NewTextBlockObject("mrkdwn", "*<fakeLink.toYourApp.com|Use Case Catalogue - CF Presentation - [June 12, 2018]>*\nThis is presentation will continue to be updated as...", false, false)
	resultFourSection := uim.NewSectionBlock(resultFourTxt, nil, uim.NewAccessory(availableOption))

	// Result Five
	resultFiveTxt := uim.NewTextBlockObject("mrkdwn", "*<fakeLink.toYourApp.com|Comprehensive Benefits Catalogue - 2019>*\nInformation about all the benfits we offer is...", false, false)
	resultFiveSection := uim.NewSectionBlock(resultFiveTxt, nil, uim.NewAccessory(availableOption))

	// Next Results Button
	// Suggestions Action
	btnTxt := uim.NewTextBlockObject("plain_text", "Next 5 Results", false, false)
	nextBtn := uim.NewButtonBlockElement("", "click_me_123", btnTxt)
	actionBlock := uim.NewActionBlock("", nextBtn)

	// Build Message with blocks created above
	msg := uim.NewBlockMessage(
		headerSection,
		divSection,
		resultOneSection,
		resultTwoSection,
		resultThreeSection,
		resultFourSection,
		resultFiveSection,
		divSection,
		actionBlock,
	)

	b, err := json.MarshalIndent(msg, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(b))

}

func unmarshalExample() {
	var msgBlocks []uim.Block

	// Append ActionBlock for marshalling
	btnTxt := uim.NewTextBlockObject("plain_text", "Add a suggestion", false, false)
	nextBtn := uim.NewButtonBlockElement("", "click_me_123", btnTxt)
	approveBtnTxt := uim.NewTextBlockObject("plain_text", "Approve", false, false)
	approveBtn := uim.NewButtonBlockElement("", "click_me_123", approveBtnTxt)
	msgBlocks = append(msgBlocks, uim.NewActionBlock("", nextBtn, approveBtn))

	// Append ContextBlock for marshalling
	profileOne := uim.NewImageBlockElement("https://api.uim.com/img/blocks/bkb_template_images/profile_1.png", "Michael Scott")
	profileTwo := uim.NewImageBlockElement("https://api.uim.com/img/blocks/bkb_template_images/profile_2.png", "Dwight Schrute")
	textBlockObj := uim.NewTextBlockObject("mrkdwn", "*<fakeLink.toHotelPage.com|Omni Royal Orleans Hotel>*\n★★★★★\n$419 per night\nRated: 8.8 - Excellent", false, false)
	msgBlocks = append(msgBlocks, uim.NewContextBlock("", []uim.MixedElement{profileOne, profileTwo, textBlockObj}...))

	// Append ImageBlock for marshalling
	msgBlocks = append(msgBlocks, uim.NewImageBlock("https://api.uim.com/img/blocks/bkb_template_images/profile_2.png", "some profile", "image-block", textBlockObj))

	// Append DividerBlock for marshalling
	msgBlocks = append(msgBlocks, uim.NewDividerBlock())

	// Append SectionBlock for marshalling
	approvalText := uim.NewTextBlockObject("mrkdwn", "*Type:*\nPaid time off\n*When:*\nAug 10-Aug 13\n*Hours:* 16.0 (2 days)\n*Remaining balance:* 32.0 hours (4 days)\n*Comments:* \"Family in town, going camping!\"", false, false)
	approvalImage := uim.NewImageBlockElement("https://api.uim.com/img/blocks/bkb_template_images/approvalsNewDevice.png", "computer thumbnail")
	msgBlocks = append(msgBlocks, uim.NewSectionBlock(approvalText, nil, uim.NewAccessory(approvalImage)), nil)

	// Build Message with blocks created above
	msg := uim.NewBlockMessage(msgBlocks...)

	b, err := json.Marshal(&msg)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(b))

	// Unmarshal message
	m := uim.Message{}
	if err := json.Unmarshal(b, &m); err != nil {
		fmt.Println(err)
		return
	}

	var respBlocks []uim.Block
	for _, block := range m.Blocks.BlockSet {
		// Need to implement a type switch to determine Block type since the
		// response from UIM could include any/all types under "blocks" key
		switch block.BlockType() {
		case uim.MBTContext:
			var respMixedElements []uim.MixedElement
			contextElements := block.(*uim.ContextBlock).ContextElements.Elements
			// Need to implement a type switch for ContextElements for same reason as Blocks
			for _, elem := range contextElements {
				switch elem.MixedElementType() {
				case uim.MixedElementImage:
					// Assert the block's type to manipulate/extract values
					imageBlockElem := elem.(*uim.ImageBlockElement)
					imageBlockElem.ImageURL = "https://api.uim.com/img/blocks/bkb_template_images/profile_1.png"
					imageBlockElem.AltText = "MichaelScott"
					respMixedElements = append(respMixedElements, imageBlockElem)
				case uim.MixedElementText:
					textBlockElem := elem.(*uim.TextBlockObject)
					textBlockElem.Text = "go go go go go"
					respMixedElements = append(respMixedElements, textBlockElem)
				}
			}
			respBlocks = append(respBlocks, uim.NewContextBlock("new block", respMixedElements...))
		case uim.MBTAction:
			actionBlock := block.(*uim.ActionBlock)
			// Need to implement a type switch for BlockElements for same reason as Blocks
			for _, elem := range actionBlock.Elements.ElementSet {
				switch elem.ElementType() {
				case uim.METImage:
					imageElem := elem.(*uim.ImageBlockElement)
					fmt.Printf("do something with image block element: %v\n", imageElem)
				case uim.METButton:
					buttonElem := elem.(*uim.ButtonBlockElement)
					fmt.Printf("do something with button block element: %v\n", buttonElem)
				case uim.METOverflow:
					overflowElem := elem.(*uim.OverflowBlockElement)
					fmt.Printf("do something with overflow block element: %v\n", overflowElem)
				case uim.METDatepicker:
					datepickerElem := elem.(*uim.DatePickerBlockElement)
					fmt.Printf("do something with datepicker block element: %v\n", datepickerElem)
				}
			}
			respBlocks = append(respBlocks, block)
		case uim.MBTImage:
			// Simply re-append the block if you want to include it in the response
			respBlocks = append(respBlocks, block)
		case uim.MBTSection:
			respBlocks = append(respBlocks, block)
		case uim.MBTDivider:
			respBlocks = append(respBlocks, block)
		}
	}

	// Build new Message with Blocks obtained/edited from callback
	respMsg := uim.NewBlockMessage(respBlocks...)

	b, err = json.Marshal(&respMsg)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(b))
}
