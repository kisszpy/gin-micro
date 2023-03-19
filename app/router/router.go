package router

import (
	"github.com/gin-gonic/gin"
	"sample/app/controller"
)

func AppRouter(e *gin.Engine) {
	group := e.Group("/api/v1")
	{
		group.GET("/hello", controller.Hello)
	}
}
