package source

import (
	"log"

	"github.com/PuerkitoBio/goquery"
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
	article []struct {
		link string
		seen bool
	}
}

// GetArticle get all article Data from Dev.To
func GetArticle(url string) []Article {
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

		data = append(data, art)
	})

	return data
}
