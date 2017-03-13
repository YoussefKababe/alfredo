package message

import (
  "encoding/json"
  "net/http"
  "dropbot/config"
  "bytes"
)

func HandleMessage(event *Event) {
  message := event.Message

  switch {
  case message.Text != "":
    handleText(event)
  case message.Attachments != nil:
    handleAttachments(event)
  }
}

func sendMessage(event *map[string]interface{}) {
  message, _ := json.Marshal(event)
  request, _ := http.NewRequest("POST", "https://graph.facebook.com/v2.6/me/messages?access_token=" + config.PageToken, bytes.NewBuffer(message))
  request.Header.Set("Content-Type", "application/json")

  client := http.Client{}
  client.Do(request)
}
