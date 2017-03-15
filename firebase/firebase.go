package firebase

import (
	"alfredo/config"
	"bytes"
	"encoding/json"
	"io/ioutil"
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

// GetUser get a user record from the Firebase database.
func GetUser(userID string) map[string]interface{} {
	request, _ := http.NewRequest("GET", "https://alfredo-b0f06.firebaseio.com/users/"+userID+".json?auth="+config.FirebaseToken, nil)
	client := http.Client{}
	resp, _ := client.Do(request)
	result, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var user map[string]interface{}
	json.Unmarshal(result, &user)
	return user
}
