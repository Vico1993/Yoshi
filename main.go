package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/vico1993/Yoshi/notification"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal("Error reading body : ", err)
		}

		// TODO : Need to find a way to hundle Http Request, and cmdline interface..
		// TODO : Need to hundle a multiple way of notification ( Messenger, text, vocale... )
		notification.GetTelegramMessage(body)
	})

	log.Println("Serving on localhost:3000")
	err := http.ListenAndServe(":3000", nil)
	log.Fatal(err)
}
