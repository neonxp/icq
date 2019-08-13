package icq

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
)

type messages struct {
	client *client
}

func newMessages(client *client) *messages {
	return &messages{client: client}
}

func (f *messages) SendText(chatID string, text string, replyMsgID []string, forwardChatID string, forwardMsgID string) (*Msg, error) {
	params := url.Values{
		"chatId": []string{chatID},
		"text":   []string{text},
	}
	if replyMsgID != nil && len(replyMsgID) > 0 {
		for _, msgID := range replyMsgID {
			params.Add("replyMsgId", msgID)
		}
	}
	if forwardChatID != "" {
		params.Set("forwardChatId", forwardChatID)
	}
	if forwardMsgID != "" {
		params.Set("forwardMsgId", forwardMsgID)
	}
	resp, err := f.client.request(
		http.MethodGet,
		"/messages/sendText",
		params,
		nil,
	)
	if err != nil {
		return nil, err
	}
	result := new(Msg)
	return result, json.NewDecoder(resp).Decode(result)
}

func (f *messages) SendExistsFile(chatID string, fileID string, caption string, replyMsgID []string, forwardChatID string, forwardMsgID string) (*Msg, error) {
	params := url.Values{
		"chatId":  []string{chatID},
		"fileId":  []string{fileID},
		"caption": []string{caption},
	}
	if replyMsgID != nil && len(replyMsgID) > 0 {
		for _, msgID := range replyMsgID {
			params.Add("replyMsgId", msgID)
		}
	}
	if forwardChatID != "" {
		params.Set("forwardChatId", forwardChatID)
	}
	if forwardMsgID != "" {
		params.Set("forwardMsgId", forwardMsgID)
	}
	resp, err := f.client.request(
		http.MethodGet,
		"/messages/sendFile",
		params,
		nil,
	)
	if err != nil {
		return nil, err
	}
	result := new(Msg)
	return result, json.NewDecoder(resp).Decode(result)
}

func (f *messages) SendFile(chatID string, fileName string, caption string, replyMsgID []string, forwardChatID string, forwardMsgID string) (*MsgLoadFile, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile("file", fileName)
	if err != nil {
		return nil, err
	}

	fh, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := fh.Close(); err != nil {
			log.Println(err)
		}
	}()
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return nil, err
	}

	if err := bodyWriter.Close(); err != nil {
		return nil, err
	}

	contentType := bodyWriter.FormDataContentType()

	params := url.Values{
		"chatId":  []string{chatID},
		"caption": []string{caption},
	}
	if replyMsgID != nil && len(replyMsgID) > 0 {
		for _, msgID := range replyMsgID {
			params.Add("replyMsgId", msgID)
		}
	}
	if forwardChatID != "" {
		params.Set("forwardChatId", forwardChatID)
	}
	if forwardMsgID != "" {
		params.Set("forwardMsgId", forwardMsgID)
	}
	resp, err := f.client.requestWithContentType(
		http.MethodPost,
		"/messages/sendFile",
		params,
		bodyBuf,
		contentType,
	)
	if err != nil {
		return nil, err
	}
	result := new(MsgLoadFile)
	return result, json.NewDecoder(resp).Decode(result)
}

func (f *messages) SendExistsVoice(chatID string, fileID string, replyMsgID []string, forwardChatID string, forwardMsgID string) (*Msg, error) {
	params := url.Values{
		"chatId": []string{chatID},
		"fileId": []string{fileID},
	}
	if replyMsgID != nil && len(replyMsgID) > 0 {
		for _, msgID := range replyMsgID {
			params.Add("replyMsgId", msgID)
		}
	}
	if forwardChatID != "" {
		params.Set("forwardChatId", forwardChatID)
	}
	if forwardMsgID != "" {
		params.Set("forwardMsgId", forwardMsgID)
	}
	resp, err := f.client.request(
		http.MethodGet,
		"/messages/sendVoice",
		params,
		nil,
	)
	if err != nil {
		return nil, err
	}
	result := new(Msg)
	return result, json.NewDecoder(resp).Decode(result)
}

func (f *messages) SendVoice(chatID string, fileName string, replyMsgID []string, forwardChatID string, forwardMsgID string) (*MsgLoadFile, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile("file", fileName)
	if err != nil {
		return nil, err
	}

	fh, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := fh.Close(); err != nil {
			log.Println(err)
		}
	}()
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return nil, err
	}

	if err := bodyWriter.Close(); err != nil {
		return nil, err
	}

	contentType := bodyWriter.FormDataContentType()

	params := url.Values{
		"chatId": []string{chatID},
	}
	if replyMsgID != nil && len(replyMsgID) > 0 {
		for _, msgID := range replyMsgID {
			params.Add("replyMsgId", msgID)
		}
	}
	if forwardChatID != "" {
		params.Set("forwardChatId", forwardChatID)
	}
	if forwardMsgID != "" {
		params.Set("forwardMsgId", forwardMsgID)
	}
	resp, err := f.client.requestWithContentType(
		http.MethodPost,
		"/messages/sendVoice",
		params,
		bodyBuf,
		contentType,
	)
	if err != nil {
		return nil, err
	}
	result := new(MsgLoadFile)
	return result, json.NewDecoder(resp).Decode(result)
}

func (f *messages) EditText(chatID string, text string, msgID string) (bool, error) {
	params := url.Values{
		"msgId":  []string{msgID},
		"chatId": []string{chatID},
		"text":   []string{text},
	}
	resp, err := f.client.request(
		http.MethodGet,
		"/messages/editText",
		params,
		nil,
	)
	if err != nil {
		return false, err
	}
	result := new(OK)
	return result.OK, json.NewDecoder(resp).Decode(result)
}

func (f *messages) DeleteMessages(chatID string, msgIDs []string) (bool, error) {
	params := url.Values{
		"chatId": []string{chatID},
	}
	if msgIDs != nil && len(msgIDs) > 0 {
		for _, msgID := range msgIDs {
			params.Add("msgId", msgID)
		}
	}
	resp, err := f.client.request(
		http.MethodGet,
		"/messages/deleteMessages",
		params,
		nil,
	)
	if err != nil {
		return false, err
	}
	result := new(OK)
	return result.OK, json.NewDecoder(resp).Decode(result)
}
