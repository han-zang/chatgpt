package chat_http

import (
	"chatgpt/logger"
	"chatgpt/typing"
	"net/http"

	"github.com/gin-gonic/gin"
)

func add_middle(mid typing.IMiddle) {
	f := func(mid typing.IMiddle) gin.HandlerFunc {
		return func(c *gin.Context) {
			// 执行前置操作
			mid.Before(rou)

			// 继续处理请求
			c.Next()

			// 执行后置操作
			mid.After(rou)
		}
	}(mid)
	rou.r.Use(f)
}

func add_event(grp typing.IRouGroup) {
	f := func(h typing.RouHandle) gin.HandlerFunc {
		return func(c *gin.Context) {
			h(rou, c)
			c.Next()
		}
	}

	g := rou.r.Group(grp.Name())
	for _, v := range grp.Events() {
		switch v.Method() {
		case http.MethodPost:
			g.POST(v.Uri(), f(v.Handle()))
		case http.MethodGet:
			g.GET(v.Uri(), f(v.Handle()))
		default:
			logger.Error("no http method %s", v.Method())
		}
	}
}
