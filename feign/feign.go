package feign

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"io"
	"net/http"
	"sample/config"
)

type Feign struct {
	NamingClient naming_client.INamingClient
}

func NewFeign(namingClient naming_client.INamingClient) *Feign {
	feign := Feign{NamingClient: namingClient}
	return &feign
}

func (f *Feign) Get(serviceName string, uri string) *RestResult {
	instances, err := f.NamingClient.SelectInstances(vo.SelectInstancesParam{
		ServiceName: serviceName,
		GroupName:   config.Cfg.Nacos.GroupName,
		HealthyOnly: true,
	})
	if err != nil {
		panic(err)
	}
	for _, instance := range instances {
		url := fmt.Sprintf("http://%s:%d%s", instance.Ip, instance.Port, uri)
		resp, e := http.Get(url)
		if e != nil {
			continue
		}
		data, e := io.ReadAll(resp.Body)
		if e != nil {
			continue
		}
		var result RestResult
		json.Unmarshal(data, &result)
		return &result
	}
	return nil
}

func Post() {

}
