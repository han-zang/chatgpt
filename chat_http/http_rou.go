package chat_http

import (
	"bytes"
	"chatgpt/typing"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (self *http_rou) CreateEvent(m string, u string, h typing.RouHandle) typing.IEvent {
	return &Event{
		method: m,
		uri:    u,
		handle: h,
	}
}

func (self *http_rou) Resp(c typing.IContext, code int, arg interface{}) {
	ctx := c.(*gin.Context)
	ctx.JSON(code, arg)
}

func (self *http_rou) Resp200(c typing.IContext, arg interface{}) {
	ctx := c.(*gin.Context)
	ctx.JSON(http.StatusOK, arg)
}

func (self *http_rou) Set(c typing.IContext, key string, arg interface{}) {
	ctx := c.(*gin.Context)
	ctx.Set(key, arg)
}

func (self *http_rou) Get(c typing.IContext, key string) (interface{}, bool) {
	ctx := c.(*gin.Context)
	return ctx.Get(key)
}

func (self *http_rou) Body(c typing.IContext, arg interface{}) error {
	ctx := c.(*gin.Context)
	return ctx.BindJSON(arg)
}

func (self *http_rou) GetRequest(c typing.IContext, key string) interface{} {
	ctx := c.(*gin.Context)
	switch key {
	case "uri":
		return ctx.Request.URL.Path
	case "code":
		return ctx.Writer.Status()
	case "body":
		body, _ := ctx.GetRawData()
		ctx.Request.Body = ioutil.NopCloser(bytes.NewReader(body))
		return string(body)
	}
	return `''`
}
