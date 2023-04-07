package chat_http

import "chatgpt/typing"

type Event struct {
	method string
	uri    string
	handle typing.RouHandle
}

func (self *Event) Method() string {
	return self.method
}

func (self *Event) Uri() string {
	return self.uri
}

func (self *Event) Handle() typing.RouHandle {
	return self.handle
}
