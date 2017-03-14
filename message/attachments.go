package message

import "dropbot/dropbox"

func handleAttachments(event *Event) {
	message := event.Message
	senderID := event.Sender.ID

	for _, attachment := range message.Attachments {
		url := attachment.Payload["url"]
		go dropbox.UploadAttachment(&url)
	}

	sendText("Your file is on its way to your Dropbox account!", senderID)
}
