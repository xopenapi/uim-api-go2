package uimtest

import (
	"testing"
	"time"

	"github.com/xopenapi/uim-api-go"
	"github.com/stretchr/testify/assert"
)

func TestRTMInfo(t *testing.T) {
	maxWait := 10 * time.Millisecond
	s := NewTestServer()
	go s.Start()

	api := uim.New("ABCDEFG", uim.OptionAPIURL(s.GetAPIURL()))
	rtm := api.NewRTM()
	go rtm.ManageConnection()
	messageChan := make(chan (*uim.ConnectedEvent), 1)
	go func() {
		for msg := range rtm.IncomingEvents {
			switch ev := msg.Data.(type) {
			case *uim.ConnectedEvent:
				messageChan <- ev
			}
		}
	}()
	select {
	case m := <-messageChan:
		assert.Equal(t, s.BotID, m.Info.User.ID, "bot id did not match")
		assert.Equal(t, s.BotName, m.Info.User.Name, "bot name did not match")
		break
	case <-time.After(maxWait):
		assert.FailNow(t, "did not get connected event in time")

	}
}

func TestRTMPing(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping timered test")
	}
	maxWait := 45 * time.Second
	s := NewTestServer()
	go s.Start()

	api := uim.New("ABCDEFG", uim.OptionAPIURL(s.GetAPIURL()))
	rtm := api.NewRTM()
	go rtm.ManageConnection()
	messageChan := make(chan (*uim.LatencyReport), 1)
	go func() {
		for msg := range rtm.IncomingEvents {
			switch ev := msg.Data.(type) {
			case *uim.LatencyReport:
				messageChan <- ev
			}
		}
	}()
	select {
	case m := <-messageChan:
		assert.NotEmpty(t, m.Value, "latency report should value a value")
		assert.True(t, m.Value > 0, "latency report should be greater than 0")
		break
	case <-time.After(maxWait):
		assert.FailNow(t, "did not get latency report in time")

	}
}

func TestRTMDirectMessage(t *testing.T) {
	maxWait := 5 * time.Second
	s := NewTestServer()
	go s.Start()

	api := uim.New("ABCDEFG", uim.OptionAPIURL(s.GetAPIURL()))
	rtm := api.NewRTM()
	go rtm.ManageConnection()
	messageChan := make(chan (*uim.MessageEvent), 1)
	go func() {
		for msg := range rtm.IncomingEvents {
			switch ev := msg.Data.(type) {
			case *uim.MessageEvent:
				messageChan <- ev
			}
		}
	}()
	s.SendDirectMessageToBot("some text")
	select {
	case m := <-messageChan:
		assert.Equal(t, defaultNonBotUserID, m.User)
		assert.Equal(t, "D024BE91L", m.Channel)
		assert.Equal(t, "some text", m.Text)
		break
	case <-time.After(maxWait):
		assert.FailNow(t, "did not get direct message in time")
	}
}

func TestRTMChannelMessage(t *testing.T) {
	maxWait := 5 * time.Second
	s := NewTestServer()
	go s.Start()

	api := uim.New("ABCDEFG", uim.OptionAPIURL(s.GetAPIURL()))
	rtm := api.NewRTM()
	go rtm.ManageConnection()
	messageChan := make(chan (*uim.MessageEvent), 1)
	go func() {
		for msg := range rtm.IncomingEvents {
			switch ev := msg.Data.(type) {
			case *uim.MessageEvent:
				messageChan <- ev
			}
		}
	}()
	s.SendMessageToChannel("#foochan", "some text")
	select {
	case m := <-messageChan:
		assert.Equal(t, "#foochan", m.Channel)
		assert.Equal(t, "some text", m.Text)
		break
	case <-time.After(maxWait):
		assert.FailNow(t, "did not get channel message in time")
	}

}
