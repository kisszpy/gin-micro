package registry

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"net"
	"sample/config"
)

type Registry struct {
}

func NewRegistry(serviceName string, port uint64) naming_client.INamingClient {
	clientConfig := constant.ClientConfig{
		NamespaceId:         config.Cfg.Nacos.NameSpaceId,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}

	// 创建服务配置
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      "localhost",
			Port:        8848,
			ContextPath: "/nacos",
		},
	}
	// Create naming client for service registration
	namingClient, err := clients.CreateNamingClient(map[string]interface{}{
		"clientConfig":  clientConfig,
		"serverConfigs": serverConfigs,
	})
	// Register Gin service to Nacos
	success, err := namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          GetLocalIPAddress(),
		Port:        port,
		ServiceName: serviceName,
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Metadata:    map[string]string{"version": "1.0"},
		ClusterName: "DEFAULT",
		GroupName:   config.Cfg.Nacos.GroupName,
		Ephemeral:   true,
	})

	if err != nil || !success {
		fmt.Println("Failed to register Gin service to Nacos:", err)
		panic("error")
	}
	fmt.Println("service register successful on nacos")
	return namingClient
}

func GetLocalIPAddress() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.String()
			}
		}
	}
	return ""
}
