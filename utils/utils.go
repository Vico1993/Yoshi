package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config is the structure of my confile file
type Config struct {
	PathToJSONFile string `json:"PathToJsonFile"`
}

// GetConfigData return all parameter of Conf.json
func GetConfigData() Config {
	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}

	return configuration
}
