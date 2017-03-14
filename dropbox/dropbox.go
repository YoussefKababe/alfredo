package dropbox

import (
	"bytes"
	"dropbot/config"
	"encoding/json"
	"fmt"
	"net/http"
)

// UploadAttachment uploads an attachment to dropbox
func UploadAttachment(url *string) {
	attachment := map[string]interface{}{
		"url":  url,
		"path": "/Apps/Alfredo/test.png",
	}

	data, _ := json.Marshal(attachment)
	request, _ := http.NewRequest("POST", "https://api.dropboxapi.com/2/files/save_url", bytes.NewBuffer(data))
	request.Header.Set("Authorization", "Bearer "+config.DropboxToken)
	request.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, _ := client.Do(request)

	fmt.Println(resp)
}
