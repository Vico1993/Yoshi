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

// GetConfigData return all parameter of Conf.json
func GetConfigData() (Config, error) {
	file, errOpeningFile := os.Open("conf.json")
	if errOpeningFile != nil {
		fmt.Println("An error append when Yoshi try to acces to your conf.json. Please Make sure the file is here.")
	}

	decoder := json.NewDecoder(file)
	configuration := Config{}
	err := decoder.Decode(&configuration)

	return configuration, err
}

// CheckConfig is here to check if the config file is here, and if all data in it has change.
func CheckConfig(configuration Config) {

}
