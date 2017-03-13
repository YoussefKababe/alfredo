package message

import (
  "encoding/json"
  "net/http"
  "dropbot/config"
  "bytes"
)

func ReceivedMessage(event *Event) {
  senderId := event.Sender.Id
  message := event.Message

  newMessage := Event{
    Message: &Message{
      Text: message.Text,
    },
    Recipient: &Person{
      Id: senderId,
    },
  }

  mm, _ := json.Marshal(newMessage)

  request, _ := http.NewRequest("POST", "https://graph.facebook.com/v2.6/me/messages?access_token=" + config.PageToken, bytes.NewBuffer(mm))
  request.Header.Set("Content-Type", "application/json")

  client := http.Client{}
  client.Do(request)
}
