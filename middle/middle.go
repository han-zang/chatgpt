package middle

import (
	"chatgpt/typing"
)

func List() []typing.IMiddle {
	return []typing.IMiddle{
		&mid_recover,
		&mid_log,
	}
}
