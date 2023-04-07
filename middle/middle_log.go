package middle

import (
	"chatgpt/typing"
)

type MiddleLog struct {
}

var mid_log MiddleLog

func (*MiddleLog) Before(c typing.IRou) {
	// logger.Info("start")
}

func (*MiddleLog) After(typing.IRou) {
	// logger.Info("over")
}
