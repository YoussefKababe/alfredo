package alfredo

import (
	"alfredo/config"
	"alfredo/dropbox"
	"alfredo/firebase"
	"alfredo/messenger"
)

// HandleMessage dfd
func HandleMessage(event *messenger.Event) {
	message := event.Message

	switch {
	case message.Text != "":
		handleText(event)
	case message.Attachments != nil:
		handleAttachments(event)
	}
}

func handleText(event *messenger.Event) {
	senderID := event.Sender.ID
	messenger.SendText("I don't know how to send text to Dropbox yet! :( try sending"+
		" me a file instead!", senderID)
}

func handleAttachments(event *messenger.Event) {
	fb := firebase.New(config.FirebaseProjectID, config.FirebaseSecret)
	db := dropbox.New(config.DropboxKey, config.DropboxSecret)
	message := event.Message
	senderID := event.Sender.ID
	user := fb.GetUser(senderID)

	if user["dropboxToken"] == nil {
		handleGetStarted(event)
		return
	}

	for _, attachment := range message.Attachments {
		go db.UploadAttachment(attachment.Payload["url"], user["dropboxToken"].(string))
	}

	messenger.SendText("Your file is on the way to your Dropbox account!", senderID)
}
