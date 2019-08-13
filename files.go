package icq

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type files struct {
	client *client
}

func newFiles(client *client) *files {
	return &files{client: client}
}

func (f *files) GetInfo(fileID string) (*FileInfo, error) {
	resp, err := f.client.request(
		http.MethodGet,
		"/chats/getInfo",
		url.Values{
			"fileId": []string{fileID},
		},
		nil,
	)
	if err != nil {
		return nil, err
	}
	result := new(FileInfo)
	return result, json.NewDecoder(resp).Decode(result)
}
