package firebase

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Firebase represents a firebase instance.
type Firebase struct {
	ProjectID string
	Secret    string
}

// New creates a new firebase instance
func New(firebaseProject, firebaseSecret string) Firebase {
	f := Firebase{
		ProjectID: firebaseProject,
		Secret:    firebaseSecret,
	}

	return f
}

//SaveUser saves a user to the Firebase database.
func (f *Firebase) SaveUser(userID, dropboxToken string) {
	user := map[string]interface{}{
		"dropboxToken": dropboxToken,
	}

	muser, _ := json.Marshal(user)

	request, _ := http.NewRequest("PUT", "https://"+f.ProjectID+".firebaseio.com/users/"+userID+".json?auth="+f.Secret, bytes.NewBuffer(muser))
	client := http.Client{}
	client.Do(request)
}

// GetUser get a user record from the Firebase database.
func (f *Firebase) GetUser(userID string) map[string]interface{} {
	request, _ := http.NewRequest("GET", "https://"+f.ProjectID+".firebaseio.com/users/"+userID+".json?auth="+f.Secret, nil)
	client := http.Client{}
	resp, _ := client.Do(request)
	result, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var user map[string]interface{}
	json.Unmarshal(result, &user)
	return user
}
