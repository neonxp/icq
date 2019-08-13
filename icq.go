package icq

type Api struct {
	Self     *self
	Chats    *chats
	Files    *files
	Messages *messages
	Events   *events
}

func NewApi(token string, apiType ApiType) *Api {
	client := newClient(token, apiType)
	return &Api{
		Self:     newSelf(client),
		Chats:    newChats(client),
		Files:    newFiles(client),
		Messages: newMessages(client),
		Events:   newEvents(client),
	}
}
