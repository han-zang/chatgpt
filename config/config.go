package config

import (
	"chatgpt/lib"

	"github.com/BurntSushi/toml"
)

type CfgOpenAi struct {
	ApiKey string `toml:"api_key"`
}

type CfgLogger struct {
	Level   string `toml:"level"`
	Console bool   `toml:"console"`
}

// type CfgEnv struct {
// 	Logger CfgLogger `toml:"logger"`
// }

type Config struct {
	OpenAi CfgOpenAi `toml:"openai"`
	Logger CfgLogger `toml:"logger"`
}

// 非线程安全的！！！！
var config Config

func Init() {
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		lib.Panic("Failed to parse TOML data: %s\n", err)
		return
	}
}

func ApiKey() string {
	return config.OpenAi.ApiKey
}

func LogLevel() string {
	if config.Logger.Console {
		return "console"
	}
	return config.Logger.Level
}
