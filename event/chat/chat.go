package chat

import (
	"chatgpt/typing"
	"sync"
)

type group_chat struct {
	lst_event []typing.IEvent
}

var e *group_chat
var once sync.Once

func Get(rou typing.IRou) *group_chat {
	once.Do(func() {
		e = &group_chat{}
		e.lst_event = append(e.lst_event, rou.CreateEvent("POST", "/talk", talk))
	})
	return e
}

func (self *group_chat) Name() string {
	return "/chat"
}

func (self *group_chat) Events() []typing.IEvent {
	return self.lst_event
}
