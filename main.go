package main

import "github.com/vico1993/Yoshi/cmd"

func main() {

	cmd.RootCmd.Execute()

	// Check Config
	// config := utils.GetConfigData()

	// tags := []string{"linux", "go", "beginners", "productivity", "php", "explainlikeimfive", "devtips", "docker", "tips", "bots"}
	// data := source.GetArticle("https://dev.to", tags)

	// if len(data) >= 1 {

	// 	cl := telegram.NewBetaClient(config.TelegramChatID, config.TelegramBotAPI)
	// 	for _, article := range data {
	// 		cl.SendTelegramMessage(article.Link+"\n"+strings.Join(article.Tags, ""), false)
	// 	}

	// 	source.UpdateArticleSent(data)
	// }
}
