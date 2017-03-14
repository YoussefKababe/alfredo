package message

import (
	"alfredo/config"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// HandleMessage decides what to do with a new event depending on the message type.
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
	request, _ := http.NewRequest("POST", "https://graph.facebook.com/v2.6/me/messages?access_token="+config.PageToken, bytes.NewBuffer(message))
	request.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	client.Do(request)
}

// SetGetStartedButton sets the get started button in messenger
func SetGetStartedButton() {
	data, _ := json.Marshal(map[string]interface{}{
		"get_started": map[string]string{
			"payload": "GET_STARTED_PAYLOAD",
		},
	})

	request, _ := http.NewRequest("POST", "https://graph.facebook.com/v2.6/me/messenger_profile?access_token="+config.PageToken, bytes.NewBuffer(data))
	request.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, _ := client.Do(request)

	fmt.Println(resp)
}
