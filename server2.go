package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"sample/app/router"
)

func main() {
	clientConfig := constant.ClientConfig{
		NamespaceId:         "public",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}

	// 创建服务配置
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: "localhost",
			Port:   8848,
		},
	}
	// Create naming client for service registration
	namingClient, err := clients.CreateNamingClient(map[string]interface{}{
		"clientConfig":  clientConfig,
		"serverConfigs": serverConfigs,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	// Register Gin service to Nacos
	success, err := namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          "127.0.0.1",
		Port:        8080,
		ServiceName: "GinService",
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Metadata:    map[string]string{"version": "1.0"},
		ClusterName: "DEFAULT",
		GroupName:   "DEFAULT_GROUP",
		Ephemeral:   true,
	})

	if err != nil || !success {
		fmt.Println("Failed to register Gin service to Nacos:", err)
		return
	}

	// Start Gin server
	r := gin.Default()
	router.AppRouter(r)
	r.Run(":8080")
}
