package icq

import "encoding/json"

type Bot struct {
	UserID    string `json:"userId"`    // уникальный идентификатор
	Nick      string `json:"nick"`      // уникальный ник
	FirstName string `json:"firstName"` // имя
	About     string `json:"about"`     // описание бота
	Photo     []struct {
		URL string `json:"url"` // url
	} `json:"photo"` // аватар бота
	OK bool `json:"ok"` // статус запроса
}

type Chat struct {
	InviteLink string `json:"inviteLink"`
	Public     bool   `json:"public"`
	Title      string `json:"title"`
	Group      string `json:"group"`
	OK         bool   `json:"ok"` // статус запроса
}

type Admin struct {
	UserID  string `json:"user_id"`
	Creator bool   `json:"creator"`
}

type Admins struct {
	Admins []Admin `json:"admins"`
}

type FileInfo struct {
	Type     string `json:"type"`
	Size     int    `json:"size"`
	Filename string `json:"filename"`
	URL      string `json:"url"`
}

type Msg struct {
	MsgID string `json:"msgId"`
	OK    bool   `json:"ok"` // статус запроса
}

type MsgLoadFile struct {
	FileID string `json:"fileId"`
	MsgID  string `json:"msgId"`
	OK     bool   `json:"ok"` // статус запроса
}

type User struct {
	UserID    string `json:"userId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type File struct {
	FileID string `json:"fileId"`
}

type EventType string

const (
	EventTypeDataMessage     EventType = "newMessage"
	EventTypeEditedMessage   EventType = "editedMessage"
	EventTypeDeletedMessage  EventType = "deletedMessage"
	EventTypePinnedMessage   EventType = "pinnedMessage"
	EventTypeUnpinnedMessage EventType = "unpinnedMessage"
	EventTypeNewChatMembers  EventType = "newChatMembers"
	EventTypeLeftChatMembers EventType = "leftChatMembers"
)

type EventInterface interface {
	GetEventID() int
	GetType() EventType
}

type Events struct {
	Events []EventInterface `json:"events"`
}

type RawEvents struct {
	Events []json.RawMessage `json:"events"`
}

type Event struct {
	EventID int       `json:"eventId"`
	Type    EventType `json:"type"`
}

func (e Event) GetEventID() int {
	return e.EventID
}

func (e Event) GetType() EventType {
	return e.Type
}

type EventDataMessage struct {
	Event
	Payload struct {
		MsgID string `json:"msgId"`
		Chat  struct {
			ChatID string `json:"chatId"`
			Type   string `json:"type"`
			Title  string `json:"title"`
		} `json:"chat"`
		From      User   `json:"from"`
		Timestamp int    `json:"timestamp"`
		Text      string `json:"text"`
		Parts     []AttachmentInterface
		RawParts  []json.RawMessage `json:"parts"`
	} `json:"payload"`
}

type EventEditedMessage struct {
	Event
	Payload struct {
		MsgID string `json:"msgId"`
		Chat  struct {
			ChatID string `json:"chatId"`
			Type   string `json:"type"`
			Title  string `json:"title"`
		} `json:"chat"`
		From            User   `json:"from"`
		Timestamp       int    `json:"timestamp"`
		Text            string `json:"text"`
		EditedTimestamp string `json:"editedTimestamp"`
	} `json:"payload"`
}

type EventDeletedMessage struct {
	Event
	Payload struct {
		MsgID string `json:"msgId"`
		Chat  struct {
			ChatID string `json:"chatId"`
			Type   string `json:"type"`
			Title  string `json:"title"`
		} `json:"chat"`
		Timestamp int `json:"timestamp"`
	} `json:"payload"`
}

type EventPinnedMessage struct {
	Event
	Payload struct {
		MsgID string `json:"msgId"`
		Chat  struct {
			ChatID string `json:"chatId"`
			Type   string `json:"type"`
			Title  string `json:"title"`
		} `json:"chat"`
		From      User   `json:"from"`
		Timestamp int    `json:"timestamp"`
		Text      string `json:"text"`
	} `json:"payload"`
}

type EventUnpinnedMessage struct {
	Event
	Payload struct {
		MsgID string `json:"msgId"`
		Chat  struct {
			ChatID string `json:"chatId"`
			Type   string `json:"type"`
			Title  string `json:"title"`
		} `json:"chat"`
		Timestamp int `json:"timestamp"`
	} `json:"payload"`
}

type EventNewChatMembers struct {
	Event
	Payload struct {
		MsgID string `json:"msgId"`
		Chat  struct {
			ChatID     string `json:"chatId"`
			NewMembers []User `json:"newMembers"`
			AddedBy    User   `json:"addedBy"`
		} `json:"chat"`
		Timestamp int `json:"timestamp"`
	} `json:"payload"`
}

type EventLeftChatMembers struct {
	Event
	Payload struct {
		MsgID string `json:"msgId"`
		Chat  struct {
			ChatID      string `json:"chatId"`
			LeftMembers []User `json:"leftMembers"`
			RemovedBy   User   `json:"removedBy"`
		} `json:"chat"`
		Timestamp int `json:"timestamp"`
	} `json:"payload"`
}

type AttachmentType string

const (
	AttachmentTypeSticker AttachmentType = "sticker"
	AttachmentTypeMention AttachmentType = "mention"
	AttachmentTypeVoice   AttachmentType = "voice"
	AttachmentTypeFile    AttachmentType = "file"
	AttachmentTypeForward AttachmentType = "forward"
	AttachmentTypeReply   AttachmentType = "reply"
)

type AttachmentInterface interface {
	GetType() AttachmentType
}

type Attachment struct {
	Type AttachmentType `json:"type"`
}

func (a Attachment) GetType() AttachmentType {
	return a.Type
}

type AttachmentSticker struct {
	Attachment
	Payload File `json:"payload"`
}

type AttachmentMention struct {
	Attachment
	Payload User `json:"payload"`
}

type AttachmentVoice struct {
	Attachment
	Payload File `json:"payload"`
}

type AttachmentFile struct {
	Attachment
	Payload struct {
		FileID  string             `json:"fileId"`
		Type    AttachmentFileType `json:"type"`
		Caption string             `json:"caption"`
	} `json:"payload"`
}

type AttachmentFileType string

const (
	AttachmentFileTypeImage AttachmentFileType = "image"
	AttachmentFileTypeAudio AttachmentFileType = "audio"
	AttachmentFileTypeVideo AttachmentFileType = "video"
)

type AttachmentForward struct {
	Attachment
	Payload struct {
		Message string `json:"message"`
	} `json:"payload"`
}

type AttachmentReply struct {
	Attachment
	Payload struct {
		Message string `json:"message"`
	} `json:"payload"`
}

type Error struct {
	OK          bool   `json:"ok"`
	Description string `json:"description"`
}

type OK struct {
	OK bool `json:"ok"`
}

type ChatAction string

const (
	ChatActionLooking ChatAction = "looking"
	ChatActionTyping  ChatAction = "typing"
)
