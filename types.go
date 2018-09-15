package icq

type Response struct {
	Response struct {
		StatusCode int              `json:"statusCode"`
		StatusText string           `json:"statusText"`
		RequestId  string           `json:"requestId"`
		Data       *MessageResponse `json:"data"`
	} `json:"response"`
}

type MessageResponse struct {
	SubCode struct {
		Error int `json:"error"`
	} `json:"subCode"`
	MessageID        string `json:"msgId"`
	HistoryMessageID int64  `json:"histMsgId"`
	State            string `json:"state"`
}
