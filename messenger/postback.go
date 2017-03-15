package messenger

import (
	"alfredo/firebase"
	"time"
)

func handleGetStarted(event *Event) {
	senderID := event.Sender.ID
	firebase.SaveUser(senderID, "")

	sendText("Hello there! I'm Alfredo, and I'll help you save your important"+
		" files to Dropbox without leaving messenger!", senderID)

	sendSenderAction(senderID, "typing_on")
	time.Sleep(time.Second * 2)

	sendDropoxAuthLink(senderID)
}
