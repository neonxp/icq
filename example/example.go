package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-icq/icq"
)

func main() {
	// New API object
	b := icq.NewAPI(os.Getenv("ICQ_TOKEN"))

	// Send message
	r, err := b.SendMessage("429950", "Hello, world!")
	if err != nil {
		panic(err)
	}
	log.Println(r.State)

	// Send file
	f, err := os.Open("./example/icq.png")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	file, err := b.UploadFile("icq.png", f)
	if err != nil {
		panic(err)
	}
	b.SendMessage("429950", file)

	// Webhook usage
	updates := make(chan icq.Update)
	errors := make(chan error)
	http.HandleFunc("/webhook", b.GetWebhookHandler(updates, errors))
	go http.ListenAndServe(":8080", nil)
	for {
		select {
		case u := <-updates:
			log.Println("Incomming message", u)
			b.SendMessage(u.Update.From.ID, fmt.Sprintf("You sent me: %s", u.Update.Text))
		case err := <-errors:
			log.Fatalln(err)
		}
	}
}
