package dropbox

import (
	"bytes"
	"alfredo/config"
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strings"
)

// UploadAttachment uploads an attachment to dropbox
func UploadAttachment(url string) {
	attachment := map[string]interface{}{
		"url":  url,
		"path": "/" + path.Base(url)[0:strings.Index(path.Base(url), "?")],
	}

	data, _ := json.Marshal(attachment)
	request, _ := http.NewRequest("POST", "https://api.dropboxapi.com/2/files/save_url", bytes.NewBuffer(data))
	request.Header.Set("Authorization", "Bearer "+config.DropboxToken)
	request.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, _ := client.Do(request)

	fmt.Println(resp)
}
