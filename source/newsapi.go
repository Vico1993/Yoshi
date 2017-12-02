package source

// import (
// 	"encoding/json"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"strings"
// )

// const newsAPIKey = "582df2397b4e4b61b61ee36a4f7e9356"

// type Article struct {
// 	Author      string `json:"author"`
// 	Title       string `json:"title"`
// 	Description string `json:"description"`
// 	URL         string `json:"url"`
// 	Publish     string `json:"publishedAt"`
// }

// type News struct {
// 	Status   bool      `json:"ok"`
// 	Source   string    `json:"source"`
// 	Sort     string    `json:"sortBy"`
// 	Articles []Article `json:"articles"`
// }

// func AskNewsApi(source string, sort string) News {

// 	var url = "https://newsapi.org/v1/articles?source=" + source + "&sortBy=" + sort + "&apiKey=" + newsAPIKey
// 	// Build the request
// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		log.Fatal("NewRequest: ", err)
// 	}

// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	// For control over HTTP client headers,
// 	// redirect policy, and other settings,
// 	// create a Client
// 	// A Client is an HTTP client
// 	client := &http.Client{}

// 	// Send the request via a client
// 	// Do sends an HTTP request and
// 	// returns an HTTP response
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Fatal("Do: ", err)
// 	}

// 	// Callers should close resp.Body
// 	// when done reading from it
// 	// Defer the closing of the body
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatal("erreur ReadAll: ", err)
// 	}
// 	result := string(body)

// 	var newsReturn News

// 	json.NewDecoder(strings.NewReader(result)).Decode(&newsReturn)

// 	return newsReturn
// }
