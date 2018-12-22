package icq

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// UploadFile to ICQ servers and returns URL to file
func (a *API) UploadFile(fileName string, r io.Reader) (*FileResponse, error) {
	v := url.Values{}
	v.Set("aimsid", a.token)
	v.Set("filename", fileName)
	req, err := http.NewRequest(http.MethodPost, a.baseUrl+"/im/sendFile?"+v.Encode(), r)
	if err != nil {
		return nil, err
	}
	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	file := struct {
		Data FileResponse `json:"data"`
	}{}
	if err := json.Unmarshal(b, &file); err != nil {
		return nil, err
	}
	return &file.Data, nil
}
