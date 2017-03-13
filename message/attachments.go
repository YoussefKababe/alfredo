package message

func HandleAttachments(event *Event) {
  senderId := event.Sender.Id
  message := event.Message

  for _, attachment := range message.Attachments {
    newMessage := map[string]interface{}{
      "message": map[string]interface{}{
        "attachment": attachment,
      },
      "recipient": map[string]interface{}{
        "id": senderId,
      },
    }

    go SendMessage(&newMessage)
  }
}
