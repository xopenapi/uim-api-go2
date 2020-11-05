package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/xopenapi/uim-api-go"
)

var api = uim.New("YOUR_TOKEN")
var signingSecret = "YOUR_SIGNING_SECRET"

// You can open a dialog with a user interaction. (like pushing buttons, slash commands ...)
// https://api.uim.com/surfaces/modals
// https://api.uim.com/interactivity/entry-points
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

	// Read request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("[ERROR] Fail to read request body: %v", err)
		return
	}

	// Verify signing secret
	sv, err := uim.NewSecretsVerifier(r.Header, signingSecret)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		log.Printf("[ERROR] Fail to verify SigningSecret: %v", err)
		return
	}
	sv.Write(body)
	if err := sv.Ensure(); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		log.Printf("[ERROR] Fail to verify SigningSecret: %v", err)
		return
	}

	// Parse request body
	str, _ := url.QueryUnescape(string(body))
	str = strings.Replace(str, "payload=", "", 1)
	var message uim.InteractionCallback
	if err := json.Unmarshal([]byte(str), &message); err != nil {
		log.Printf("[ERROR] Fail to unmarchal json: %v", err)
		return
	}

	switch message.Type {
	case uim.InteractionTypeInteractionMessage:
		// Make new dialog components and open a dialog.
		// Component-Text
		textInput := uim.NewTextInput("TextSample", "Sample label - Text", "Default value")

		// Component-TextArea
		textareaInput := uim.NewTextAreaInput("TexaAreaSample", "Sample label - TextArea", "Default value")

		// Component-Select menu
		option1 := uim.DialogSelectOption{
			Label: "Display name 1",
			Value: "Inner value 1",
		}
		option2 := uim.DialogSelectOption{
			Label: "Display name 2",
			Value: "Inner value 2",
		}
		options := []uim.DialogSelectOption{option1, option2}
		selectInput := uim.NewStaticSelectDialogInput("SelectSample", "Sample label - Select", options)

		// Open a dialog
		elements := []uim.DialogElement{
			textInput,
			textareaInput,
			selectInput,
		}
		dialog := uim.Dialog{
			CallbackID:  "Callback_ID",
			Title:       "Dialog title",
			SubmitLabel: "Submit",
			Elements:    elements,
		}
		api.OpenDialog(message.TriggerID, dialog)

	case uim.InteractionTypeDialogSubmission:
		// Receive a notification of a dialog submission
		log.Printf("Successfully receive a dialog submission.")
	}
}
