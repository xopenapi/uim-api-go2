package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/xopenapi/uim-api-go"
	"github.com/xopenapi/uim-api-go/uimevents"
)

// You more than likely want your "Bot User OAuth Access Token" which starts with "xoxb-"
var api = uim.New("TOKEN")

func main() {
	http.HandleFunc("/events-endpoint", func(w http.ResponseWriter, r *http.Request) {
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		body := buf.String()
		eventsAPIEvent, e := uimevents.ParseEvent(json.RawMessage(body), uimevents.OptionVerifyToken(&uimevents.TokenComparator{VerificationToken: "TOKEN"}))
		if e != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if eventsAPIEvent.Type == uimevents.URLVerification {
			var r *uimevents.ChallengeResponse
			err := json.Unmarshal([]byte(body), &r)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			w.Header().Set("Content-Type", "text")
			w.Write([]byte(r.Challenge))
		}
		if eventsAPIEvent.Type == uimevents.CallbackEvent {
			innerEvent := eventsAPIEvent.InnerEvent
			switch ev := innerEvent.Data.(type) {
			case *uimevents.AppMentionEvent:
				api.PostMessage(ev.Channel, uim.MsgOptionText("Yes, hello.", false))
			}
		}
	})
	fmt.Println("[INFO] Server listening")
	http.ListenAndServe(":3000", nil)
}
