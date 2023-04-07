package chat

import (
	"chatgpt/config"
	"chatgpt/logger"
	"chatgpt/typing"
	"context"
	"net/http"

	"github.com/sashabaranov/go-openai"
)

func talk(rou typing.IRou, ctx typing.IContext) {
	client := openai.NewClient(config.ApiKey())
	// resp, err := client.CreateCompletion(context.Background(), openai.CompletionRequest{
	// 	Model:  openai.GPT3TextAda001,
	// 	Prompt: "My name is zz.\nQ:what is your name?\nA:",
	// })
	// if err != nil {
	// 	logger.Error("chat complete error: %v\n", err)
	// 	rou.Resp(ctx, http.StatusBadRequest, nil)
	// 	return
	// }
	// rou.Resp(ctx, http.StatusOK, resp)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Hello!",
				},
				{
					Role:    openai.ChatMessageRoleAssistant,
					Content: "Hi there! How can I assist you today?",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Tell me the last thing you said.",
				},
			},
		},
	)
	if err != nil {
		logger.Error("chat complete error: %v\n", err)
		rou.Resp(ctx, http.StatusBadRequest, nil)
		return
	}
	rou.Resp(ctx, http.StatusOK, resp)
}
