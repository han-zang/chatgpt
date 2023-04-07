package chat_http

import (
	"chatgpt/typing"
	"sync"

	"github.com/gin-gonic/gin"
)

type http_rou struct {
	r *gin.Engine
}

var rou *http_rou
var once sync.Once

func Get() typing.IRou {
	return rou
}

func Init() {
	once.Do(func() {
		if rou == nil {
			rou = &http_rou{
				r: gin.New(),
			}
			gin.SetMode(gin.ReleaseMode)

		}
	})
}

func Start(lst_mid []typing.IMiddle, lst_grp []typing.IRouGroup) {
	for _, v := range lst_mid {
		add_middle(v)
	}
	for _, v := range lst_grp {
		add_event(v)
	}
	rou.r.Run("127.0.0.1:8888")
}
