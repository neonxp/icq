package icq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"time"
)

type ApiType int

const (
	ICQ ApiType = iota
	Agent
)

var servers = map[ApiType]string{
	ICQ:   "https://api.icq.net/bot/v1/",
	Agent: "https://agent.mail.ru/bot/v1/",
}

type client struct {
	token   string
	apiType ApiType
	client  http.Client
}

func newClient(token string, apiType ApiType) *client {
	return &client{token: token, apiType: apiType, client: http.Client{Timeout: 30 * time.Second}}
}

func (c *client) request(method string, methodPath string, query url.Values, body *bytes.Buffer) (io.Reader, error) {
	return c.requestWithContentType(method, methodPath, query, body, "")
}

func (c *client) requestWithContentType(method string, methodPath string, query url.Values, body *bytes.Buffer, contentType string) (io.Reader, error) {
	query.Set("token", c.token)
	u, err := url.Parse(servers[c.apiType])
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, methodPath)
	u.RawQuery = query.Encode()

	req, err := http.NewRequest(method, u.String(), nil)
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	if body != nil {
		rc := ioutil.NopCloser(body)
		req.Body = rc
	}

	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		errObj := new(Error)
		err = json.NewDecoder(resp.Body).Decode(errObj)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("ok=%v message=%s", errObj.OK, errObj.Description)
	}
	return resp.Body, err
}
