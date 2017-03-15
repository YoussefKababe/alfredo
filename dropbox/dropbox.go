package dropbox

import (
	"alfredo/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
)

// UploadAttachment uploads an attachment to dropbox
func UploadAttachment(url, token string) {
	attachment := map[string]interface{}{
		"url":  url,
		"path": "/" + path.Base(url)[0:strings.Index(path.Base(url), "?")],
	}

	data, _ := json.Marshal(attachment)
	request, _ := http.NewRequest("POST", "https://api.dropboxapi.com/2/files/save_url", bytes.NewBuffer(data))
	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, _ := client.Do(request)

	fmt.Println(resp)
}

// GetAuthToken converts a Dropbox API code into and auth token.
func GetAuthToken(code string) string {
	data := url.Values{}
	data.Set("code", code)
	data.Add("grant_type", "authorization_code")
	data.Add("client_id", config.DropboxKey)
	data.Add("client_secret", config.DropboxSecret)
	data.Add("redirect_uri", "https://dropbot.localtunnel.me/mdropbox")

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
