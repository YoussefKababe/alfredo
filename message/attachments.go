package message

import "dropbot/dropbox"

func handleAttachments(event *Event) {
	message := event.Message
	senderID := event.Sender.ID

	for _, attachment := range message.Attachments {
		go dropbox.UploadAttachment(attachment.Payload["url"])
	}

	sendText("Your file is on its way to your Dropbox account!", senderID)
}
