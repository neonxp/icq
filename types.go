package icq

type Response struct {
	Response struct {
		StatusCode int              `json:"statusCode"`
		StatusText string           `json:"statusText"`
		RequestId  string           `json:"requestId"`
		Data       *MessageResponse `json:"data"`
	} `json:"response"`
}
type ParseType string

const (
	ParseURL         ParseType = "url"
	ParseFilesharing           = "filesharing"
)

type Message struct {
	To       string
	Text     string
	Mentions []string
	Parse    []ParseType
}

type MessageResponse struct {
	SubCode struct {
		Error int `json:"error"`
	} `json:"subCode"`
	MessageID        string `json:"msgId"`
	HistoryMessageID int64  `json:"histMsgId"`
	State            string `json:"state"`
}

type FileResponse struct {
	StaticUrl     string `json:"static_url"`
	MimeType      string `json:"mime"`
	SnapID        string `json:"snapId"`
	TtlID         string `json:"ttl_id"`
	IsPreviewable int    `json:"is_previewable"`
	FileID        string `json:"fileid"`
	FileSize      int    `json:"filesize"`
	FileName      string `json:"filename"`
	ContentID     string `json:"content_id"`
}

type WebhookRequest struct {
	Token   string   `json:"aimsid"`
	Updates []Update `json:"update"`
}

type Update struct {
	Update struct {
		Chat Chat   `json:"chat"`
		Date int    `json:"date"`
		From User   `json:"from"`
		Text string `json:"text"`
	} `json:"update"`
	UpdateID int `json:"update_id"`
}

type Chat struct {
	ID string `json:"id"`
}

type User struct {
	ID           string `json:"id"`
	LanguageCode string `json:"language_code"`
}
