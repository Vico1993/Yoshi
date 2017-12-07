package main

import (
	"strings"

	telegram "github.com/vico1993/Telegram-Client"
	"github.com/vico1993/Yoshi/source"
)

func main() {
	tags := []string{"linux", "go", "beginners", "productivity", "php", "explainlikeimfive", "devtips", "docker", "tips", "bots"}
	data := source.GetArticle("https://dev.to", tags)

	if len(data) >= 1 {
		cl := telegram.NewBetaClient("359897077", "429433832:AAHhjwe5-IQXoXTU0gduQuFDsQnilA7RKLU")
		cl.SendTelegramMessage("Voici la Front Page de dev.to", true)
		for _, article := range data {
			cl.SendTelegramMessage(article.Link+"\n"+strings.Join(article.Tags, ""), false)
		}

		source.UpdateArticleSent(data)
	}
}
