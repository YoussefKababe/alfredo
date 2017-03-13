package message

func handleText(event *Event) {
	senderId := event.Sender.Id
	message := event.Message

	newMessage := map[string]interface{}{
		"message": map[string]interface{}{
			"text": message.Text,
		},
		"recipient": map[string]interface{}{
			"id": senderId,
		},
	}

	go sendMessage(&newMessage)
}
