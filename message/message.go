package message

import (
	"alfredo/config"
	"alfredo/dropbox"
	"alfredo/firebase"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
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

// HandlePostback decides what to do with a new event depending on the postback payload.
func HandlePostback(event *Event) {
	payload := event.Postback.Payload

	switch payload {
	case "GET_STARTED_PAYLOAD":
		handleGetStarted(event)
	}
}

func sendMessage(event *map[string]interface{}) {
	message, _ := json.Marshal(event)
	request, _ := http.NewRequest("POST", "https://graph.facebook.com/v2.6/me/messages?access_token="+config.PageToken, bytes.NewBuffer(message))
	request.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	client.Do(request)
}

func sendSenderAction(recipientID, action string) {
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
				"text": "Hi {{user_full_name}}! Do you often receive a lot of files" +
					" on messenger? Forward them to me and I will put them right into your Dropbox!",
			},
		},
	}

	updateMessengerProfile(data)
}

// LinkDropbox links a Dropbox account to a user.
func LinkDropbox(c echo.Context) error {
	code := c.QueryParam("code")
	userID := c.QueryParam("state")

	token := dropbox.GetAuthToken(code)

	firebase.SaveUser(userID, token)
	sendText("Super! I keep getting told I look like a cat but, I'm not really"+
		" good at anything cats can do :( I'm only good at saving files to Dropbox!"+
		" Forward all the important files you have on messenger to me and I'll"+
		" instantly put them into your Dropbox!", userID)
	return c.String(200, "You're Dropbox account was successfully linked! You"+
		" can close this tab and go back to messenger.")
}
