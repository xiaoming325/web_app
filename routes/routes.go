package routes

import (
	"web_app/logger"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()                                      //没有任何中间件
	r.Use(logger.GinLogger(), logger.GinRecovery(true)) //添加自定义的中间件
	return r
}
