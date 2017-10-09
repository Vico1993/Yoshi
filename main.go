package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"

	"github.com/vico1993/Yoshi/api"
	"github.com/vico1993/Yoshi/notification"
)

func stringIndexToVar(text string) (string, []string) {
	tmp := strings.Split(text, " ")
	return tmp[0], tmp[1:]
}

func hundleCommand(cmd string) {

	var params []string

	// Need to hundle sentences into commands
	if !strings.ContainsAny(cmd, "/") {
		// TODO : Make some scenario
	} else {
		if strings.ContainsAny(cmd, " ") {
			cmd, params = stringIndexToVar(cmd)
		}
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
		var source string

		if len(params) > 0 {
			source = params[0]
		} else {
			source = "bbc-news"
		}

		notification.SendTelegramMessage("Alors, voici quelque news de "+source+" ", true)

		var newsReturn = api.AskNewsApi(source, "top")

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
			println("Error reading body : ", err)
		}

		// TODO : Need to find a way to hundle Http Request, and cmdline interface..
		// TODO : Need to hundle a multiple way of notification ( Messenger, text, vocale... )
		var cmd = notification.GetTelegramCommand(body)

		if cmd != "" {
			hundleCommand(cmd)
		}
	})

	notification.SendTelegramMessage("INFO : Yoshi on Port : 3000", true)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		notification.SendTelegramMessage("ERROR : Impossible de démarrer mon Serveur.. Yoshi Out..", true)
	}
}
