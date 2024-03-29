package messenger

import (
	"alfredo/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func sendMessage(event *map[string]interface{}) {
	message, _ := json.Marshal(event)
	request, _ := http.NewRequest("POST", "https://graph.facebook.com/v2.6/me/messages?access_token="+config.PageToken, bytes.NewBuffer(message))
	request.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	client.Do(request)
}

// SendSenderAction sends user action to messenger
func SendSenderAction(recipientID, action string) {
	data := map[string]interface{}{
		"recipient": map[string]string{
			"id": recipientID,
		},
		"sender_action": action,
	}

	sendMessage(&data)
}

func updateMessengerProfile(data map[string]interface{}) {
	mdata, _ := json.Marshal(data)
	request, _ := http.NewRequest("POST", "https://graph.facebook.com/v2.6/me/messenger_profile?access_token="+config.PageToken, bytes.NewBuffer(mdata))
	request.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, _ := client.Do(request)

	if resp.StatusCode >= 400 {
		defer resp.Body.Close()
		content, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("updateMessengerProfile:", string(content))
	}
}

// SetGetStartedButton sets the get started button in messenger
func SetGetStartedButton() {
	data := map[string]interface{}{
		"get_started": map[string]string{
			"payload": "GET_STARTED_PAYLOAD",
		},
	}

	updateMessengerProfile(data)
}

// SetGreetingText sets the greeting text in messenger
func SetGreetingText() {
	data := map[string]interface{}{
		"greeting": []map[string]string{
			map[string]string{
				"locale": "default",
				"text": "Hi {{user_first_name}}! Do you often receive a lot of files" +
					" on messenger? Forward them to me and I will save them instantly to your Dropbox!",
			},
		},
	}

	updateMessengerProfile(data)
}
