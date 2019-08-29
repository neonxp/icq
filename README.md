# Проект архивирован потому что появилась официальная библиотека https://github.com/mail-ru-im/bot-golang

# ICQ Bot Api Go

[![Sourcegraph](https://sourcegraph.com/github.com/go-icq/icq/-/badge.svg?style=flat-square)](https://sourcegraph.com/github.com/go-icq/icq?badge)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/go-icq/icq)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-icq/icq?style=flat-square)](https://goreportcard.com/report/github.com/go-icq/icq)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/go-icq/icq/master/LICENSE)

Основана на новом Bot Api (https://icq.com/botapi/)

Реализованы все методы и соответствуют документации.
## Методы
```go
api.Events.Get(ctx context.Context) <-chan EventInterface
api.Self.Get() (*Bot, error)
api.Messages.SendText(chatID string, text string, replyMsgID []string, forwardChatID string, forwardMsgID string) (*Msg, error)
api.Messages.SendExistsFile(chatID string, fileID string, caption string, replyMsgID []string, forwardChatID string, forwardMsgID string) (*Msg, error)
api.Messages.SendFile(chatID string, fileName string, caption string, replyMsgID []string, forwardChatID string, forwardMsgID string) (*MsgLoadFile, error)
api.Messages.SendExistsVoice(chatID string, fileID string, replyMsgID []string, forwardChatID string, forwardMsgID string) (*Msg, error)
api.Messages.SendVoice(chatID string, fileName string, replyMsgID []string, forwardChatID string, forwardMsgID string) (*MsgLoadFile, error)
api.Messages.EditText(chatID string, text string, msgID string) (bool, error)
api.Messages.DeleteMessages(chatID string, msgIDs []string) (bool, error)
api.Chats.SendActions(chatID string, actions []ChatAction) (bool, error)
api.Chats.GetInfo(chatID string) (*Chat, error)
api.Chats.GetAdmins(chatID string) (*Admins, error)
api.Files.GetInfo(fileID string) (*FileInfo, error)
```

Типы можно увидеть в http://godoc.org/github.com/go-icq/icq

## Пример

```go
package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/go-icq/icq"
)

func main() {
	// Инициализация
	b := icq.NewApi(os.Getenv("ICQ_TOKEN"), icq.ICQ) // or icq.Agent

	// Получение информации о боте
	log.Println(b.Self.Get())

	// Отправка сообщения
	resultSend, err := b.Messages.SendText("429950", "Привет!", nil, "", "")
	if err != nil {
		log.Fatal(err)
	}

	// Отправка файла
	resultFile, err := b.Messages.SendFile("429950", "./example/example.jpg", "коржик", []string{resultSend.MsgID}, "", "")
	if err != nil {
		log.Fatal(err)
	}

	// Отправка существующего файла по ID
	_, err = b.Messages.SendExistsFile("429950", resultFile.FileID, "Существующий файл", nil, "", "")
	if err != nil {
		log.Fatal(err)
	}

	// Редактирование сообщения
	_, err = b.Messages.EditText("429950", "Новый текст", resultSend.MsgID)
	if err != nil {
		log.Fatal(err)
	}

	// Будем слушать эвенты 5 минут. При закрытии контекста перестает работать цикл получения событий. В реальном мире контекст надо будет закрывать по сигналу ОС
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Minute)
	for ev := range b.Events.Get(ctx) {
		switch ev := ev.(type) {
		case *icq.EventDataMessage:
			b.Messages.SendText(ev.Payload.Chat.ChatID, "Echo: "+ev.Payload.Text, []string{ev.Payload.MsgID}, "", "")
		default:
			log.Println(ev)
		}
	}
}
```

## Автор

Александр NeonXP Кирюхин  <a.kiryukhin@mail.ru>
