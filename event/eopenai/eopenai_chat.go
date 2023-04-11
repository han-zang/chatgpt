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
	status := msg.StatusFail
	resp := &msg.MsgTalkResp{}
	for {
		// var req msg.MsgTalkReq
		req := &msg.MsgTalkReq{}
		if err := rou.Body(ctx, req); err != nil {
			logger.Error("openai chat parse body err %s", err.Error())
			break
		}
		chat_req := openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
		}
		for i := 0; i < len(req.Messages); i++ {
			v := req.Messages[i]
			if r, succ := role(v.Role); succ {
				chat_req.Messages = append(chat_req.Messages, openai.ChatCompletionMessage{
					Role:    r,
					Content: v.Context,
				})
			}
		}
		client := openai.NewClient(config.ApiKey())

		chat_resp, err := client.CreateChatCompletion(
			context.Background(),
			chat_req,
		)
		if err != nil {
			logger.Error("openai chat req err %s", err.Error())
			break
		}
		resp.Context = chat_resp.Choices[0].Message.Content
		status = msg.StatusOk
		break
	}
	resp.Status = status
	rou.Resp200(ctx, resp)
}

func role(role msg.MsgRoleType) (string, bool) {
	switch role {
	case msg.RoleType_User:
		return openai.ChatMessageRoleUser, true
	case msg.RoleType_Ass:
		return openai.ChatMessageRoleAssistant, true
	}
	return "", false
}
