package icq

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type chats struct {
	client *client
}

func newChats(client *client) *chats {
	return &chats{client: client}
}

func (s *chats) SendActions(chatID string, actions []ChatAction) (bool, error) {
	acts := []string{}
	for _, act := range actions {
		acts = append(acts, string(act))
	}
	resp, err := s.client.request(
		http.MethodGet,
		"/chats/sendActions",
		url.Values{
			"chatId":  []string{chatID},
			"actions": acts,
		},
		nil,
	)
	if err != nil {
		return false, err
	}
	result := new(OK)
	return result.OK, json.NewDecoder(resp).Decode(result)
}

func (s *chats) GetInfo(chatID string) (*Chat, error) {
	resp, err := s.client.request(
		http.MethodGet,
		"/chats/getInfo",
		url.Values{
			"chatId": []string{chatID},
		},
		nil,
	)
	if err != nil {
		return nil, err
	}
	result := new(Chat)
	return result, json.NewDecoder(resp).Decode(result)
}

func (s *chats) GetAdmins(chatID string) (*Admins, error) {
	resp, err := s.client.request(
		http.MethodGet,
		"/chats/getAdmins",
		url.Values{
			"chatId": []string{chatID},
		},
		nil,
	)
	if err != nil {
		return nil, err
	}
	result := new(Admins)
	return result, json.NewDecoder(resp).Decode(result)
}
