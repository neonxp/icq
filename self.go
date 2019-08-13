package icq

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type self struct {
	client *client
}

func newSelf(client *client) *self {
	return &self{client: client}
}

func (s *self) Get() (*Bot, error) {
	resp, err := s.client.request(http.MethodGet, "/self/get", url.Values{}, nil)
	if err != nil {
		return nil, err
	}
	result := new(Bot)
	return result, json.NewDecoder(resp).Decode(result)
}
