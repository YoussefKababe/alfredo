package message

import (
  "encoding/json"
  "net/http"
  "dropbot/config"
  "bytes"
  "fmt"
)

func HandleMessage(event *Event) {
  message := event.Message

  switch {
  case message.Text != "":
    HandleText(event)
  case message.Attachments != nil:
    HandleAttachments(event)
  }
}

func SendMessage(event *map[string]interface{}) {
  message, _ := json.Marshal(event)
  request, _ := http.NewRequest("POST", "https://graph.facebook.com/v2.6/me/messages?access_token=" + config.PageToken, bytes.NewBuffer(message))
  request.Header.Set("Content-Type", "application/json")

  client := http.Client{}
  resp, _ := client.Do(request)

  fmt.Println(resp)
}
