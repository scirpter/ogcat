package common

import (
	"encoding/json"
	"os"
)

const CONFIG_FILE_NAME = "config.json"

type Config struct {
	DiscordTokens   []string `json:"discordTokens"`
	DiscordPassword string   `json:"discordPassword"`
}

func ReadJSONFile() *Config {
	file, err := os.Open(CONFIG_FILE_NAME)
	if err != nil {
		return nil
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	config := &Config{}
	err = decoder.Decode(config)
	if err != nil {
		return nil
	}
	return config
}

func WriteToJSONFile(config *Config) error {
	file, err := os.OpenFile(CONFIG_FILE_NAME, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(config)
	if err != nil {
		return err
	}
	return nil
}
