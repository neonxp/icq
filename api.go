package icq

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// HTTP Client interface
type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

// API
type API struct {
	token   string
	baseUrl string
	client  Doer
}

// NewAPI constructor of API object
func NewAPI(token string) *API {
	return &API{
		token:   token,
		baseUrl: "https://botapi.icq.net",
		client:  http.DefaultClient,
	}
}

// SendMessage with `message` text to `to` participant
func (a *API) SendMessage(message Message) (*MessageResponse, error) {
	parse, _ := json.Marshal(message.Parse)
	v := url.Values{}
	v.Set("aimsid", a.token)
	v.Set("r", strconv.FormatInt(time.Now().Unix(), 10))
	v.Set("t", message.To)
	v.Set("message", message.Text)
	v.Set("mentions", strings.Join(message.Mentions, ","))
	if len(message.Parse) > 0 {
		v.Set("parse", string(parse))
	}
	b, err := a.send("/im/sendIM", v)
	if err != nil {
		return nil, err
	}
	r := &Response{}
	if err := json.Unmarshal(b, r); err != nil {
		return nil, err
	}
	if r.Response.StatusCode != 200 {
		return nil, fmt.Errorf("failed to send message: %s", r.Response.StatusText)
	}
	return r.Response.Data, nil
}

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

// GetWebhookHandler returns http.HandleFunc that parses webhooks
func (a *API) GetWebhookHandler(cu chan<- Update, e chan<- error) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		if r.Method != http.MethodPost {
			e <- fmt.Errorf("incorrect method: %s", r.Method)
			return
		}
		wr := &WebhookRequest{}
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			e <- err
			return
		}
		if err := json.Unmarshal(b, wr); err != nil {
			e <- err
			return
		}
		for _, u := range wr.Updates {
			cu <- u
		}
	}
}

func (a *API) send(path string, v url.Values) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, a.baseUrl+path, strings.NewReader(v.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return b, fmt.Errorf("ICQ API error. Code=%d Message=%s", resp.StatusCode, string(b))
	}
	return b, nil
}
