package middle

import (
	"chatgpt/logger"
	"chatgpt/typing"
)

type MiddleLog struct {
}

var mid_log MiddleLog

func (*MiddleLog) Before(c typing.IRou, ctx typing.IContext) {
	logger.Info("http conneted. uri:%s code:%d body:%s", c.GetRequest(ctx, "uri"), c.GetRequest(ctx, "code"), c.GetRequest(ctx, "body"))
}

func (*MiddleLog) After(c typing.IRou, ctx typing.IContext) {
	// logger.Info("http conneted. uri:%s code:%d body:%s", c.GetRequest(ctx, "uri"), c.GetRequest(ctx, "code"), c.GetRequest(ctx, "resp"))
}
