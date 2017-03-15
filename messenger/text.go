package messenger

func handleText(event *Event) {
	senderID := event.Sender.ID
	message := event.Message
	sendText(message.Text, senderID)
}

func sendText(message string, recipientID string) {
	newMessage := map[string]interface{}{
		"message": map[string]interface{}{
			"text": message,
		},
		"recipient": map[string]interface{}{
			"id": recipientID,
		},
	}

	go sendMessage(&newMessage)
}
