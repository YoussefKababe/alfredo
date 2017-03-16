package messenger

// SendText sends a text message/
func SendText(message string, recipientID string) {
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
