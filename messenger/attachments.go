package messenger

import "alfredo/config"

// SendDropoxAuthLink sends a button to link Dropbox.
func SendDropoxAuthLink(recipientID string) {
	newMessage := map[string]interface{}{
		"message": map[string]interface{}{
			"attachment": map[string]interface{}{
				"type": "template",
				"payload": map[string]interface{}{
					"template_type": "button",
					"text": "Before I can help you do that, let's link your" +
						" Dropbox account first!",
					"buttons": []map[string]string{
						map[string]string{
							"type": "web_url",
							"url": "https://www.dropbox.com/oauth2/authorize" +
								"?client_id=b2ooejf291z2tex&response_type=code" +
								"&redirect_uri=" + config.DropboxRedirect +
								"&state=" + recipientID,
							"title": "Click to link your Dropbox!",
						},
					},
				},
			},
		},
		"recipient": map[string]interface{}{
			"id": recipientID,
		},
	}

	sendMessage(&newMessage)
}
