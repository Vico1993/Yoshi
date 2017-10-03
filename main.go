package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		SendMessage("Hey je reçois quelque chose..", false)
	})

	log.Println("Serving on localhost:3000")
	err := http.ListenAndServe(":3000", nil)
	log.Fatal(err)
}

const chatTelegramID = "359897077"
const botAPI = "429433832:AAHhjwe5-IQXoXTU0gduQuFDsQnilA7RKLU"

//

type SendStruct struct {
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

func SendMessage(text string, notification bool) bool {
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
	var newrecord SendStruct

	json.NewDecoder(strings.NewReader(result)).Decode(&newrecord)
	if newrecord.Ok {
		return true
	}

	return false
}
