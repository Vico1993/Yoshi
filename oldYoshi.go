package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"

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

		if params[0] == "list" {
			notification.SendTelegramMessage("Alors, voici la liste des sources possible...  ", true)
			sources := []string{"abc-news-au-api", "al-jazeera-english-api", "ars-technica-api", "associated-press-api", "bbc-news-api", "bbc-sport-api", "bloomberg-api", "breitbart-news-api", "business-insider-api", "business-insider-uk-api", "buzzfeed-api", "cnbc-api", "cnn-api", "daily-mail-api", "engadget-api", "entertainment-weekly-api", "espn-api", "espn-cric-info-api", "financial-times-api", "football-italia-api", "fortune-api", "four-four-two-api", "fox-sports-api", "google-news-api", "hacker-news-api", "ign-api", "independent-api", "mashable-api", "metro-api", "mirror-api", "mtv-news-api", "mtv-news-uk-api", "national-geographic-api", "new-scientist-api", "newsweek-api", "new-york-magazine-api", "nfl-news-api", "polygon-api", "recode-api", "reddit-r-all-api", "reuters-api", "talksport-api", "techcrunch-api", "techradar-api", "the-economist-api", "the-guardian-au-api", "the-guardian-uk-api", "the-hindu-api", "the-huffington-post-api", "the-lad-bible-api", "the-new-york-times-api", "the-next-web-api", "the-sport-bible-api", "the-telegraph-api", "the-times-of-india-api", "the-verge-api", "the-wall-street-journal-api", "the-washington-post-api", "time-api", "usa-today-api"}

			for _, src := range sources {
				notification.SendTelegramMessage(src, false)
			}
		}

		if len(params) > 0 {
			source = params[0]
		} else {
			source = "bbc-news"
		}

		notification.SendTelegramMessage("Alors, voici quelque news de "+source+" ", true)

		var newsReturn = source.AskNewsApi(source, "top")

		for _, artcl := range newsReturn.Articles {
			var message = "*" + artcl.Title + "* \n" + artcl.Description + "\n" + artcl.URL
			notification.SendTelegramMessage(message, false)
		}
	default:
		notification.SendTelegramMessage("Je suis toujours en apprentissage.. je n'es pas compris.", true)
	}
}

func mainOld() {

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
