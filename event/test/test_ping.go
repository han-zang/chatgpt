package test

import (
	"chatgpt/typing"
	"net/http"
)

func ping(rou typing.IRou, ctx typing.IContext) {
	rou.Resp(ctx, http.StatusOK, "PONG")
}
