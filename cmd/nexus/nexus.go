package main

import (
	"flag"
	"fmt"
	"github.com/Zheaoli/nexus/config"
)

func main() {
	configPath := flag.String("c", "", "config file path")
	flag.Parse()
	if *configPath == "" {
		fmt.Printf("config file needed")
		return
	}
	configData, _ := config.ParseConfigFile(*configPath)
	fmt.Println(configData)
}
