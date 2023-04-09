package eopenai

import (
	"chatgpt/config"
	"chatgpt/logger"
	"chatgpt/msg"
	"chatgpt/typing"
	"context"

	"github.com/sashabaranov/go-openai"
)

func chat(rou typing.IRou, ctx typing.IContext) {
	var err error
	status := msg.StatusFail
	resp := &msg.MsgTalkResp{}
	for {
		var req *msg.MsgTalkReq
		if err = rou.Body(ctx, req); err != nil {
			break
		}
		chat_req := openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
		}
		for i := len(req.Context) - 1; i >= 0; i-- {
			v := req.Context[i]
			chat_req.Messages = append(chat_req.Messages, openai.ChatCompletionMessage{
				Role:    v.Role,
				Content: v.Content,
			})
		}

		client := openai.NewClient(config.ApiKey())
		_, err := client.CreateChatCompletion(
			context.Background(),
			chat_req,
		)
		if err != nil {
			break
		}
		status = msg.StatusOk
	}

	resp.Status = status
	rou.Resp200(ctx, resp)
	if err != nil {
		logger.Error("openai chat err %v %v", status, err)
	}
}
