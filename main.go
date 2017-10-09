package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/vico1993/Yoshi/api"
	"github.com/vico1993/Yoshi/notification"
)

func hundleCommand(cmd string) {

	// Need to hundle sentences into commands
	if !strings.ContainsAny(cmd, "/") {

	}

	switch cmd {
	case "/docker":
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
	case "/news":

		var newsReturn = api.AskNewsApi("recode", "top")

		for _, artcl := range newsReturn.Articles {
			var message = "*" + artcl.Title + "* \n" + artcl.Description + "\n" + artcl.URL
			notification.SendTelegramMessage(message, false)
		}
	default:
		notification.SendTelegramMessage("Je suis toujours en apprentissage.. je n'es pas compris.", true)
	}
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal("Error reading body : ", err)
		}

		// TODO : Need to find a way to hundle Http Request, and cmdline interface..
		// TODO : Need to hundle a multiple way of notification ( Messenger, text, vocale... )
		var cmd = notification.GetTelegramCommand(body)

		hundleCommand(cmd)
	})

	log.Println("Serving on localhost:3000")
	err := http.ListenAndServe(":3000", nil)
	log.Fatal(err)
}
