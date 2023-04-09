package test

import (
	"chatgpt/typing"
)

func ping(rou typing.IRou, ctx typing.IContext) {
	rou.Resp200(ctx, "PONG")
}
