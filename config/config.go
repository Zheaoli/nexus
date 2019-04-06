package config

import (
	"github.com/pelletier/go-toml"
	"io/ioutil"
)

type NexusConfig struct {
	Wechat        *WechatConfig    `toml:"wechat"`
	Server        *ServerConfig    `toml:"server"`
	Logging       *LoggingConfig   `toml:"logging"`
	WebhookConfig []*WebhookConfig `toml:"webhook"`
}

var URLMap = map[string]*WebhookConfig{}
var ConfigData *NexusConfig

func ParseConfigFile(configPath string) error {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}
	var config NexusConfig
	if err := toml.Unmarshal(data, &config); err != nil {
		return err
	}
	ConfigData = &config
	return nil
}
