package notification

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os/exec"
	"strings"
)

type sendStruct struct {
	Ok     bool `json:"ok"`
	Result struct {
		MessageID int `json:"message_id"`
		From      struct {
			ID        int    `json:"id"`
			FirstName string `json:"first_name"`
			Username  string `json:"username"`
		} `json:"from"`
		Chat struct {
			ID        int    `json:"id"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Username  string `json:"username"`
			Type      string `json:"type"`
		} `json:"chat"`
		Date int    `json:"date"`
		Text string `json:"text"`
	} `json:"result"`
}

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

const chatTelegramID = "359897077"
const botAPI = "429433832:AAHhjwe5-IQXoXTU0gduQuFDsQnilA7RKLU"

func hundleCommand(cmd string) {

	// Direct command
	if strings.ContainsAny(cmd, "/") {

		switch cmd {
		case "/docker":
			out, err := exec.Command("docker", "ps", "--format", `{{.RunningFor}}:{{.Names}}`).Output()
			if err != nil {
				fmt.Println(fmt.Sprint(err) + ": " + string(out))
			}

			for _, txt := range strings.Split(string(out), "\n") {
				dockerPs := strings.Split(txt, ":")
				if dockerPs[0] != "" {
					SendTelegramMessage("Le container : "+dockerPs[1]+" run depuis : "+dockerPs[0], true)
				}
			}
		default:
			SendTelegramMessage("Je suis toujours en apprentissage.. je n'es pas compris.", true)
		}
	} else {
		// Context command
		// TODO : Do context scenario
	}
}

func GetTelegramMessage(body []byte) {
	data := &returnStruct{}
	jsonErr := json.Unmarshal([]byte(string(body)), data)
	if jsonErr != nil {
		log.Fatal("Error json Unmarshal : ", jsonErr)
	}

	if data.Message.From.IsBot == true {
		log.Fatal("C'est un bot...")
	}

	var command = data.Message.Text

	hundleCommand(command)
}

func SendTelegramMessage(text string, notification bool) bool {
	var URL *url.URL
	URL, err := url.Parse("https://api.telegram.org/bot" + botAPI + "/sendMessage")
	if err != nil {
		panic("boom")
	}

	parameters := url.Values{}
	parameters.Add("chat_id", chatTelegramID)
	parameters.Add("parse_mode", "markdown")
	parameters.Add("text", text)
	if !notification {
		parameters.Add("disable_notification", "true")
	}
	URL.RawQuery = parameters.Encode()

	// Build the request
	req, err := http.NewRequest("GET", URL.String(), nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("erreur ReadAll: ", err)
	}
	var result = string(body)
	var newrecord sendStruct

	json.NewDecoder(strings.NewReader(result)).Decode(&newrecord)
	if newrecord.Ok {
		return true
	}

	return false
}
