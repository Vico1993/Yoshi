package main

import (
	"strings"

	"github.com/vico1993/Yoshi/notification"
	"github.com/vico1993/Yoshi/source"
)

func main() {
	data := source.GetArticle("https://dev.to")

	notification.SendTelegramMessage("Voici la Front Page de dev.to", true)
	for _, article := range data {
		notification.SendTelegramMessage(article.Link+"\n"+strings.Join(article.Tags, " - "), false)
		break
	}
}
