package alfredo

import (
	"alfredo/config"
	"alfredo/firebase"
	"alfredo/messenger"
	"time"
)

// HandlePostback decides what to do with a new event depending on the postback payload.
func HandlePostback(event *messenger.Event) {
	payload := event.Postback.Payload

	switch payload {
	case "GET_STARTED_PAYLOAD":
		handleGetStarted(event)
	}
}

func handleGetStarted(event *messenger.Event) {
	fb := firebase.New(config.FirebaseProjectID, config.FirebaseSecret)
	senderID := event.Sender.ID
	fb.SaveUser(senderID, "")

	messenger.SendText("Hello there! I'm Alfredo, and I'll help you save your important"+
		" files to Dropbox without leaving messenger!", senderID)

	messenger.SendSenderAction(senderID, "typing_on")
	time.Sleep(time.Second * 2)

	messenger.SendDropoxAuthLink(senderID)
}
