package main

import (
	"github.com/go-icq/icq"
	"log"
	"os"
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
}
