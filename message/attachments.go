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

func sendDropoxAuthLink(recipientID string) {
	newMessage := map[string]interface{}{
		"message": map[string]interface{}{
			"attachment": map[string]interface{}{
				"type": "template",
				"payload": map[string]interface{}{
					"template_type": "button",
					"text": "Before I can help you do that, you have to link me to" +
						" your Dropbox account first!",
					"buttons": []map[string]string{
						map[string]string{
							"type": "web_url",
							"url": "https://www.dropbox.com/oauth2/authorize" +
								"?client_id=b2ooejf291z2tex&response_type=token" +
								"&redirect_uri=https://dropbot.localtunnel.me/dropbox" +
								"&state=" + recipientID,
							"title": "Link my Dropbox",
						},
					},
				},
			},
		},
		"recipient": map[string]interface{}{
			"id": recipientID,
		},
	}

	sendMessage(&newMessage)
}
