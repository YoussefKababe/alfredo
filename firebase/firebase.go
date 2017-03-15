package firebase

import (
	"alfredo/config"
	"bytes"
	"encoding/json"
	"net/http"
)

//SaveUser saves a user to the Firebase database.
func SaveUser(userID, dropboxToken string) {
	user := map[string]interface{}{
		userID: map[string]string{
			"dropboxToken": dropboxToken,
		},
	}

	muser, _ := json.Marshal(user)

	request, _ := http.NewRequest("PUT", "https://alfredo-b0f06.firebaseio.com/users.json?auth="+config.FirebaseToken, bytes.NewBuffer(muser))
	client := http.Client{}
	client.Do(request)
}
