package chat_http

import (
	"chatgpt/logger"
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
			gin.SetMode(gin.ReleaseMode)
			rou = &http_rou{
				r: gin.New(),
			}
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
	url := "127.0.0.1:8888"
	logger.Info("http server start: %s", url)
	rou.r.Run(url)
}
