package message

import "alfredo/dropbox"

func handleAttachments(event *Event) {
	message := event.Message
	senderID := event.Sender.ID

	for _, attachment := range message.Attachments {
		go dropbox.UploadAttachment(attachment.Payload["url"])
	}

	sendText("Your files are on their way to your Dropbox account!", senderID)
}
