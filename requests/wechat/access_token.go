package wechat

import (
	"encoding/json"
	"github.com/Zheaoli/nexus/config"
)

func GetAccessToken() error {
	config.ConfigData.Wechat.Service.Lock.Lock()
	defer config.ConfigData.Wechat.Service.Lock.Unlock()
	params := map[string]interface{}{
		"corpid":     config.ConfigData.Wechat.Service.CorpID,
		"corpsecret": config.ConfigData.Wechat.Service.Corpsecret,
	}
	message, err := GetRequest(WECHAT_SERVER_URL+"/gettoken", params)
	if err != nil {
		return err
	}
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(message), &data); err != nil {
		return nil
	}
	if val, ok := data["errcode"]; ok {
		if val.(float64) == 0 {
			config.ConfigData.Wechat.Service.UpdateAccessToken(data["access_token"].(string))
			return nil
		} else {
			return err
		}
	} else {
		return err
	}
}

