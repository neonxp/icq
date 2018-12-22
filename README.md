# ICQ Bot API

## Installation

Go get: `go get gopkg.in/icq.v2`

Go mod / Go dep: `import "gopkg.in/icq.v2"`


## Working

Methods:

* SendMessage
* UploadFile
* FetchEvents

Webhooks workds but not recommends

## Example

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"gopkg.in/icq.v2"
)

func main() {
	// New API object
	b := icq.NewAPI(os.Getenv("ICQ_TOKEN"))

	ctx, cancel := context.WithCancel(context.Background())

	ch := make(chan interface{}) // Events channel
	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, os.Interrupt)
	signal.Notify(osSignal, os.Kill)

	go b.FetchEvents(ctx, ch) // Events fetch loop

	for {
		select {
		case e := <-ch:
			handleEvent(b, e)
		case <-osSignal:
			cancel()
			break
		}
	}
}

func handleEvent(b *icq.API, event interface{}) {
	switch event.(type) {
	case *icq.IMEvent:
		message := event.(*icq.IMEvent)
		if err := handleMessage(b, message); err != nil {
			b.SendMessage(icq.Message{
				To:   message.Data.Source.AimID,
				Text: "Message process fail",
			})
		}
	default:
		log.Printf("%#v", event)
	}
}

func handleMessage(b *icq.API, message *icq.IMEvent) error {
	cmd, ok := icq.ParseCommand(message)
	if !ok {
		return nil
	}
	_, err := b.SendMessage(icq.Message{
		To:   cmd.From,
		Text: fmt.Sprintf("Command: %s, Arguments: %v", cmd.Command, cmd.Arguments),
	})
	return err
}
```
