package main

import (
	"chatgpt/chat_http"
	"chatgpt/config"
	"chatgpt/event"
	"chatgpt/event/test"
	"chatgpt/logger"
	"chatgpt/middle"
)

func main() {
	config.Init()
	test.TestSteam()
	return

	logger.Init()
	chat_http.Init()

	chat_http.Start(middle.List(), event.List(chat_http.Get()))

	logger.Info("server stop")
}
