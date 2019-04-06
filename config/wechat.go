package config

import "sync"

type WechatConfig struct {
	Service *WechatServiceConfig `toml:"Service"`
}

type WechatServiceConfig struct {
	corpID      string       `toml:"corpid"`
	corpsecret  string       `toml:"corpsecret"`
	AccessToken string       `toml:"-"`
	Lock        sync.RWMutex `toml:"-"`
}

func (service *WechatServiceConfig) UpdateAccessToken(accessToken string) {
	service.Lock.Lock()
	defer service.Lock.Unlock()
	service.AccessToken = accessToken
}
