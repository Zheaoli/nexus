package main

import (
	"flag"
	"fmt"
	"github.com/Zheaoli/nexus/api"
	"github.com/Zheaoli/nexus/config"
	"github.com/Zheaoli/nexus/requests/wechat"
	"github.com/gin-gonic/gin"
)
func InitGroup() {
	err := wechat.GetAccessToken()
	if err != nil {
		panic(err)
	}
	for _, value := range config.ConfigData.WebhookConfig {
		config.URLMap[value.URL] = value
	}
	wechat.InitGroup()
}


func main() {
	configPath := flag.String("c", "", "config file path")
	flag.Parse()
	if *configPath == "" {
		fmt.Printf("config file needed")
		return
	}
	err := config.ParseConfigFile(*configPath)
	if err != nil {
		panic(err)
	}
	InitGroup()
	router := gin.Default()
	api.APIInit(router)
	router.Run(":8088")
}
