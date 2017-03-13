package message

func handleText(event *Event) {
	senderID := event.Sender.ID
	message := event.Message

	newMessage := map[string]interface{}{
		"message": map[string]interface{}{
			"text": message.Text,
		},
		"recipient": map[string]interface{}{
			"id": senderID,
		},
	}

	go sendMessage(&newMessage)
}
