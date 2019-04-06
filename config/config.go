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

var URLMap = map[string]interface{}{}

func ParseConfigFile(configPath string) (*NexusConfig, error) {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	var config NexusConfig
	if err := toml.Unmarshal(data, &config); err != nil {
		return nil, nil
	}
	return &config, nil
}
