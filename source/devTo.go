package source

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/vico1993/Yoshi/utils"
)

const (
	baseURL = "https://dev.to"
)

// Article Structure des article trouvé sur Dev.to
type Article struct {
	Title string
	Link  string
	Tags  []string
}

// ArticleSent Json envoyé en Telegram
type ArticleSent struct {
	Link string `json:"link"`
	Seen bool   `json:"seen"`
}

type listArticl struct {
	Collection []ArticleSent
}

// GetArticle get all article Data from Dev.To
func GetArticle(url string, hastags []string) []Article {
	var data []Article
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".articles-list .single-article").Each(func(index int, item *goquery.Selection) {
		var art Article

		// Get Link
		linkTag := item.Find(".index-article-link")
		link, _ := linkTag.Attr("href")
		art.Link = baseURL + link

		// Get Title
		art.Title = linkTag.Find("h3").Text()

		// Get Tags
		tagsNode := item.Find(".tags")
		tagsNode.Each(func(i int, item *goquery.Selection) {
			art.Tags = append(art.Tags, item.Find("span").Text())
		})

		if articleGoTag(art.Tags, hastags) {
			// If not sent append to return
			if !alreadySent(art.Link) {
				data = append(data, art)
			}
		}
	})

	return data
}

// UpdateArticleSent write new article sent
func UpdateArticleSent(data []Article) {
	dataSent := getArticleSent()

	for _, d := range data {
		newStruct := ArticleSent{
			d.Link,
			true,
		}

		dataSent = append(dataSent, newStruct)
	}

	out, err := json.Marshal(dataSent)
	if err != nil {
		panic(err)
	}

	config := utils.GetConfigData()

	file, err := os.OpenFile(config.Path+"/send/devTo.json", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("File does not exists or cannot be created")
		os.Exit(1)
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	fmt.Fprintf(w, string(out))

	w.Flush()
}

func alreadySent(link string) bool {
	ret := false
	dataSent := getArticleSent()

	for _, a := range dataSent {
		if a.Link == link {
			ret = true
			break
		}
	}

	return ret
}

func getArticleSent() []ArticleSent {
	config := utils.GetConfigData()
	data, err := ioutil.ReadFile(config.Path + "/send/devTo.json")
	if err != nil {
		kill := utils.Kill{"Can't read/find the file devTo.json, Please check your file.. ", err}
		utils.KillYoshi(kill)
	}

	var dataSent []ArticleSent
	lerr := json.Unmarshal(data, &dataSent)
	if lerr != nil {
		kill := utils.Kill{"Can't parse your JsonFile.. Please check it ! ( https://jsonlint.com ) ", lerr}
		utils.KillYoshi(kill)
	}

	return dataSent
}

func articleGoTag(artTag []string, tagWanted []string) bool {

	artTag = strings.Split(artTag[0], "#")
	res := false
	for _, at := range artTag {
		for _, tw := range tagWanted {
			if at == tw {
				res = true
				break
			}
		}
		if res {
			break
		}
	}

	return res
}
