package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/vico1993/Yoshi/notification"
)

type returnStruct struct {
	UpdateID int `json:"update_id"`
	Message  struct {
		MessageID int `json:"message_id"`
		From      struct {
			ID           int    `json:"id"`
			IsBot        bool   `json:"is_bot"`
			FirstName    string `json:"first_name"`
			LastName     string `json:"last_name"`
			LanguageCode string `json:"language_code"`
		} `json:"from"`
		Chat struct {
			ID        int    `json:"id"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Type      string `json:"type"`
		} `json:"chat"`
		Date int    `json:"date"`
		Text string `json:"text"`
	} `json:"message"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal("Error reading body : ", err)
		}

		data := &returnStruct{}
		jsonErr := json.Unmarshal([]byte(string(body)), data)
		if jsonErr != nil {
			log.Fatal("Error json Unmarshal : ", jsonErr)
		}

		if data.Message.From.IsBot == true {
			log.Fatal("C'est un bot...")
		}

		var command = data.Message.Text

		if command == "/docker" {
			out, err := exec.Command("docker", "ps", "--format", `{{.RunningFor}}:{{.Names}}`).Output()
			if err != nil {
				fmt.Println(fmt.Sprint(err) + ": " + string(out))
			}

			for _, txt := range strings.Split(string(out), "\n") {
				dockerPs := strings.Split(txt, ":")
				if dockerPs[0] != "" {
					notification.SendTelegramMessage("Le container : "+dockerPs[1]+" run depuis : "+dockerPs[0], true)
				}
			}
		} else {
			notification.SendTelegramMessage("Je suis toujours en apprentissage.. je n'es pas compris.", true)
		}

		// println(&data.Message.Text)
		// println("Bonjour, pourquoi " + data.Message.From.FirstName + " " + data.Message.From.LastName + " m'envoit tu ça : " + data.Message.Text)
	})

	log.Println("Serving on localhost:3000")
	err := http.ListenAndServe(":3000", nil)
	log.Fatal(err)
}
