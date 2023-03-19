package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var Cfg *Config

type AppConfig struct {
	Name string `mapstructure:'name'`
	Port int    `mapstructure:'port'`
}
type DbConfig struct {
	DriverType string `mapstructure:'driverType'`
	Url        string `mapstructure:'url'`
	Username   int    `mapstructure:'username'`
	Password   int    `mapstructure:'password'`
}
type NacosConfig struct {
	NacosAddress string `mapstructure:'serverAddr'`
	NameSpaceId  string `mapstructure:'nameSpaceId'`
	GroupName    string `mapstructure:'groupName'`
	ClusterName  string `mapstructure:'clusterName'`
}
type Config struct {
	App        AppConfig   `mapstructure:'app'`
	Datasource DbConfig    `mapstructure:'datasource'`
	Nacos      NacosConfig `mapstructure:nacos`
}

func readFile() {
	viper.SetConfigName("bootstrap")
	viper.AddConfigPath("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("fatal error %v \n", err)
		panic(err)
	}
	if err := viper.Unmarshal(&Cfg); err != nil {
		fmt.Printf("fatal error %v \n", err)
		panic(err)
	}
}
func init() {
	readFile()
}
