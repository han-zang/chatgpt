package middle

import (
	"chatgpt/logger"
	"chatgpt/typing"
	"runtime/debug"
)

type MiddleRecover struct {
}

var mid_recover MiddleRecover

func (*MiddleRecover) Before(c typing.IRou) {
	defer func() {
		if r := recover(); r != nil {
			// 错误处理逻辑
			logger.Error("core dump %s", string(debug.Stack()))
		}
	}()
}

func (*MiddleRecover) After(typing.IRou) {
}
