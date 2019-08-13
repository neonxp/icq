package icq

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type events struct {
	client *client
}

func newEvents(client *client) *events {
	return &events{client: client}
}

func (e *events) Get(ctx context.Context) <-chan EventInterface {
	ch := make(chan EventInterface)
	go func() {
		lastEvent := 0
		for {
			if ctx.Err() != nil {
				close(ch)
				return
			}
			events, err := e.getEvents(lastEvent)
			if err != nil {
				log.Println(err)
				<-time.After(5 * time.Second) // Retry after 5 seconds
				continue
			}
			for _, e := range events.Events {
				ch <- e
				lastEvent = e.GetEventID()
			}
		}
	}()
	return ch
}

func (e *events) getEvents(lastEvent int) (*Events, error) {
	resp, err := e.client.request(
		http.MethodGet,
		"/events/get",
		url.Values{
			"lastEventId": []string{strconv.Itoa(lastEvent)},
			"pollTime":    []string{"30"},
		},
		nil)
	if err != nil {
		return nil, err
	}
	tempResult := new(RawEvents)
	if err := json.NewDecoder(resp).Decode(tempResult); err != nil {
		return nil, err
	}
	result := new(Events)
	for _, e := range tempResult.Events {
		tempEvent := new(Event)
		if err := json.Unmarshal(e, tempEvent); err != nil {
			return nil, err
		}
		var ev EventInterface
		switch tempEvent.GetType() {
		case EventTypeDataMessage:
			ev = new(EventDataMessage)
		case EventTypeEditedMessage:
			ev = new(EventEditedMessage)
		case EventTypeDeletedMessage:
			ev = new(EventDeletedMessage)
		case EventTypePinnedMessage:
			ev = new(EventPinnedMessage)
		case EventTypeUnpinnedMessage:
			ev = new(EventUnpinnedMessage)
		case EventTypeNewChatMembers:
			ev = new(EventNewChatMembers)
		case EventTypeLeftChatMembers:
			ev = new(EventLeftChatMembers)
		}
		if err := json.Unmarshal(e, ev); err != nil {
			return nil, err
		}
		switch ev := ev.(type) {
		case *EventDataMessage:
			for _, ea := range ev.Payload.RawParts {
				tempAttachment := new(Attachment)
				if err := json.Unmarshal(ea, tempAttachment); err != nil {
					return nil, err
				}
				var eav AttachmentInterface
				switch tempAttachment.Type {
				case AttachmentTypeSticker:
					eav = new(AttachmentSticker)
				case AttachmentTypeMention:
					eav = new(AttachmentMention)
				case AttachmentTypeVoice:
					eav = new(AttachmentVoice)
				case AttachmentTypeFile:
					eav = new(AttachmentFile)
				case AttachmentTypeForward:
					eav = new(AttachmentForward)
				case AttachmentTypeReply:
					eav = new(AttachmentReply)
				}
				if err := json.Unmarshal(ea, eav); err != nil {
					return nil, err
				}
				ev.Payload.Parts = append(ev.Payload.Parts, eav)
			}
		}
		result.Events = append(result.Events, ev)
	}
	return result, nil
}
