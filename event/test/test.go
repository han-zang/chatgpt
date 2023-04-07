package test

import (
	"chatgpt/typing"
	"net/http"
	"sync"
)

type group_test struct {
	lst_event []typing.IEvent
}

var e *group_test
var once sync.Once

func Get(rou typing.IRou) *group_test {
	once.Do(func() {
		e = &group_test{}
		e.lst_event = append(e.lst_event, rou.CreateEvent(http.MethodGet, "/ping", ping))
	})
	return e
}

func (self *group_test) Name() string {
	return "/test"
}

func (self *group_test) Events() []typing.IEvent {
	return self.lst_event
}
