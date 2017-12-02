package source

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

const (
	baseURL = "https://dev.to"
)

// GetArticle get all article Data from Dev.To
func GetArticle() {
	doc, err := goquery.NewDocument(baseURL)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".articles-list .single-article").Each(func(index int, item *goquery.Selection) {
		var tags string
		// Get Link
		linkTag := item.Find(".index-article-link")
		link, _ := linkTag.Attr("href")
		link = baseURL + link

		// Get Title
		title := linkTag.Find("h3").Text()

		// Get Tags
		tagsNode := item.Find(".tags")
		tagsNode.Each(func(i int, item *goquery.Selection) {
			tags += item.Find("span").Text() + " "
		})

		fmt.Printf("Post #%d: %s - %s\nTag : %s\n\n", index, title, link, tags)
	})
}
