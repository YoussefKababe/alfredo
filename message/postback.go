package message

import "time"

func handleGetStarted(event *Event) {
	senderID := event.Sender.ID

	sendText("Hello there! I'm Alfredo, and I'll help you collect your important"+
		" messenger files in one place so they can be easy to find!", senderID)

	sendSenderAction(senderID, "typing_on")
	time.Sleep(time.Second * 4)

	sendDropoxAuthLink(senderID)
}
