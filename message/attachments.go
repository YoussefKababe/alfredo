package message

func handleAttachments(event *Event) {
	senderID := event.Sender.ID
	message := event.Message

	for _, attachment := range message.Attachments {
		newMessage := map[string]interface{}{
			"message": map[string]interface{}{
				"attachment": attachment,
			},
			"recipient": map[string]interface{}{
				"id": senderID,
			},
		}

		go sendMessage(&newMessage)
	}
}
