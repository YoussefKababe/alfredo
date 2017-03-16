package dropbox

import (
	"alfredo/config"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
)

// Dropbox represents a dropbox instance.
type Dropbox struct {
	Key    string
	Secret string
}

// New create a dropbox instance.
func New(dropboxKey, dropboxSecret string) Dropbox {
	d := Dropbox{
		Key:    dropboxKey,
		Secret: dropboxSecret,
	}

	return d
}

// UploadAttachment uploads an attachment to dropbox
func (dropbox *Dropbox) UploadAttachment(url, token string) {
	attachment := map[string]interface{}{
		"url":  url,
		"path": "/" + path.Base(url)[0:strings.Index(path.Base(url), "?")],
	}

	data, _ := json.Marshal(attachment)
	request, _ := http.NewRequest("POST", "https://api.dropboxapi.com/2/files/save_url", bytes.NewBuffer(data))
	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	client.Do(request)
}

// GetAuthToken converts a Dropbox API code into and auth token.
func (dropbox *Dropbox) GetAuthToken(code string) string {
	data := url.Values{}
	data.Set("code", code)
	data.Add("grant_type", "authorization_code")
	data.Add("client_id", dropbox.Key)
	data.Add("client_secret", dropbox.Secret)
	data.Add("redirect_uri", config.DropboxRedirect)

	request, _ := http.NewRequest("POST", "https://api.dropboxapi.com/oauth2/token", bytes.NewBufferString(data.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}
	resp, _ := client.Do(request)
	result, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	var content map[string]interface{}
	json.Unmarshal(result, &content)

	return content["access_token"].(string)
}
