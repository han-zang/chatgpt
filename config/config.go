package config

import (
	"chatgpt/logger"

	"github.com/BurntSushi/toml"
)

type CfgOpenAi struct {
	ApiKey string `toml:"api_key"`
}

type Config struct {
	OpenAi CfgOpenAi `toml:"openai"`
}

//	非线程安全的！！！！
var config Config

func Init() {
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		logger.Error("Failed to parse TOML data: %s", err)
		return
	}
}

func ApiKey() string {
	return config.OpenAi.ApiKey
}
