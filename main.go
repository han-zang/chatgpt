package main

import (
	"chatgpt/chat_http"
	"chatgpt/config"
	"chatgpt/event"
	"chatgpt/logger"
	"chatgpt/middle"
	"fmt"
)

func main() {
	logger.Init()
	config.Init()
	chat_http.Init()

	chat_http.Start(middle.List(), event.List(chat_http.Get()))

	fmt.Println("server stop")
}
