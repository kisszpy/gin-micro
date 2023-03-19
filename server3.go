package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sample/config"
	"sample/feign"
	"sample/registry"
	"strconv"
)

func main() {
	serverPort := config.Cfg.App.Port
	namingClient := registry.NewRegistry(config.Cfg.App.Name, uint64(serverPort))
	newFeign := feign.NewFeign(namingClient)
	// Start Gin server
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		// 服务发现调用
		data := newFeign.Get("GinService", "/ping")
		fmt.Printf("响应的数据： %v \n", data)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	port := ":" + strconv.Itoa(serverPort)
	r.Run(port)
}
