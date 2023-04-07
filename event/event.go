package event

import (
	"chatgpt/event/chat"
	"chatgpt/event/test"
	"chatgpt/typing"
)

func List(rou typing.IRou) []typing.IRouGroup {
	return []typing.IRouGroup{
		chat.Get(rou),
		test.Get(rou),
	}
}
