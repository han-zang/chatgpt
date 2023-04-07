package chat_http

import (
	"chatgpt/typing"

	"github.com/gin-gonic/gin"
)

func (self *http_rou) Resp(c typing.IContext, code int, arg interface{}) {
	ctx := c.(*gin.Context)
	ctx.JSON(code, arg)
}

func (self *http_rou) CreateEvent(m string, u string, h typing.RouHandle) typing.IEvent {
	return &Event{
		method: m,
		uri:    u,
		handle: h,
	}
}
