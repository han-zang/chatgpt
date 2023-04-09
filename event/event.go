package event

import (
	"chatgpt/event/eopenai"
	"chatgpt/event/test"
	"chatgpt/typing"
)

func List(rou typing.IRou) []typing.IRouGroup {
	return []typing.IRouGroup{
		eopenai.Get(rou),
		test.Get(rou),
	}
}
