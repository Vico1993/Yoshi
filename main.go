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

		notification.GetTelegramMessage(body)

		// println(&data.Message.Text)
		// println("Bonjour, pourquoi " + data.Message.From.FirstName + " " + data.Message.From.LastName + " m'envoit tu ça : " + data.Message.Text)
	})

	log.Println("Serving on localhost:3000")
	err := http.ListenAndServe(":3000", nil)
	log.Fatal(err)
}
