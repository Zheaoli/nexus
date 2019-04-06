package config

import "sync"

type WechatConfig struct {
	Service *WechatServiceConfig `toml:"Service"`
}

type WechatServiceConfig struct {
	CorpID      string       `toml:"corpid"`
	Corpsecret  string       `toml:"Corpsecret"`
	AccessToken string       `toml:"-"`
	Lock        sync.RWMutex `toml:"-"`
}

func (service *WechatServiceConfig) UpdateAccessToken(accessToken string) {
	service.AccessToken = accessToken
}
