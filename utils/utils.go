package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config is the structure of my confile file
type Config struct {
	Path           string `json:"Path"`
	TelegramChatID string `json:"telegram_chat_id"`
	TelegramBotAPI string `json:"telegram_bot_api"`
}

// Kill Struct to kill Yoshi and display an error Message
type Kill struct {
	Message string
	err     error
}

func killYoshi(k kill) {
	fmt.Println(k.Message)
	os.Exit(1)
}

// GetConfigData return all parameter of Conf.json
func GetConfigData() Config {
	file, errOpeningFile := os.Open("conf.json")
	if errOpeningFile != nil {
		killData := kill{"An error append when Yoshi try to acces to your conf.json. Please Make sure the file is here.", errOpeningFile}
		killYoshi(killData)
	}

	decoder := json.NewDecoder(file)
	configuration := Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		killData := kill{"An error append when Yoshi try to acces to your conf.json. Please check if your json is correct ( https://jsonlint.com ) .", err}
		killYoshi(killData)
	}

	if configuration.Path == "/Path/to/Go/module/Yoshi" || configuration.TelegramBotAPI == "123456789" || configuration.TelegramChatID == "my_bot_api_key:123456789" {
		killData := kill{"Please check that your enter your configuration.", err}
		killYoshi(killData)
	}

	return configuration
}

// CheckConfig is here to check if the config file is here, and if all data in it has change.
func CheckConfig(configuration Config) {

}
