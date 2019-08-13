# ICQ Bot Api Go

[![Sourcegraph](https://sourcegraph.com/github.com/go-icq/icq/-/badge.svg?style=flat-square)](https://sourcegraph.com/github.com/go-icq/icq?badge)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/go-icq/icq)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-icq/icq?style=flat-square)](https://goreportcard.com/report/github.com/go-icq/icq)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/go-icq/icq/master/LICENSE)

Основана на новом Bot Api (https://icq.com/botapi/)

Реализованы все методы и соответствуют документации.

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