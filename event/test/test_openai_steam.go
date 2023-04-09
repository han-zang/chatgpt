package test

import (
	"chatgpt/config"
	"context"
	"fmt"
	"time"

	openai "github.com/sashabaranov/go-openai"
)

func TestSteam() {
	t := time.Now().Unix()
	c := openai.NewClient(config.ApiKey())
	ctx := context.Background()
	var ls []string
	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "写一篇20字的春天文章",
			},
		},
		MaxTokens: 220,
		User:      "zh-ttttt1995",
	}
	resp, err := c.CreateChatCompletion(ctx, req)
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		return
	}
	ls = append(ls, resp.Choices[0].Message.Content)
	fmt.Println(time.Now().Unix() - t)
	fmt.Printf("Stream response: %v\n", resp)

	req = openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "我请求的上一个问题是什么",
			},
		},
		MaxTokens: 200,
		User:      "zh-ttttt1995",
	}
	resp, err = c.CreateChatCompletion(ctx, req)
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		return
	}
	ls = append(ls, resp.Choices[0].Message.Content)
	fmt.Println(time.Now().Unix() - t)
	fmt.Printf("Stream response: %v\n", resp)

}
