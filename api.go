package icq

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// HTTP Client interface
type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

// API
type API struct {
	token     string
	baseUrl   string
	client    Doer
	fetchBase string
}

// NewAPI constructor of API object
func NewAPI(token string) *API {
	return &API{
		token:   token,
		baseUrl: "https://botapi.icq.net",
		client:  http.DefaultClient,
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
