package wechat

import (
	"encoding/json"
	"github.com/Zheaoli/nexus/config"
)

func InitGroup() {
	for _, value := range config.URLMap {
		if value.Target == "wechat" {
			if err := getGroupInfo(value); err != nil {
				continue
			}
		}
	}
}

func getGroupInfo(group *config.WebhookConfig) error {
	config.ConfigData.Wechat.Service.Lock.RLock()
	defer config.ConfigData.Wechat.Service.Lock.RUnlock()
	message, err := GetRequest(WECHAT_SERVER_URL+"/appchat/get", map[string]interface{}{
		"access_token": config.ConfigData.Wechat.Service.AccessToken,
		"chatid":       group.Wechat.GroupID,
	})
	if err != nil {
		return err
	}
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(message), &data); err != nil {
		return nil
	}
	if val, ok := data["errcode"]; ok {
		if val.(float64) == 86003 {
			if err := createGroup(group.Wechat.GroupName, group.Wechat.GroupID,
				group.Wechat.GroupOwner, group.Wechat.GroupMembers); err != nil {
				return err
			}
			return nil
		} else {
			return err
		}
	} else {
		return err
	}
}
func createGroup(groupName, groupID, groupOwner string, groupMembers []string) error {
	config.ConfigData.Wechat.Service.Lock.RLock()
	defer config.ConfigData.Wechat.Service.Lock.RUnlock()
	message, err := PostRequest(WECHAT_SERVER_URL+"/appchat/create", map[string]interface{}{
		"access_token": config.ConfigData.Wechat.Service.AccessToken,
	}, map[string]interface{}{
		"name":     groupName,
		"owner":    groupOwner,
		"userlist": groupMembers,
		"chatid":   groupID,
	})
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(message), &data); err != nil {
		return nil
	}
	if val, ok := data["errcode"]; ok {
		if val.(float64) == 0 {
			return nil
		} else {
			return err
		}
	} else {
		return err
	}
}

func SendGroupMessage(groupID, message string) error {
	config.ConfigData.Wechat.Service.Lock.RLock()
	defer config.ConfigData.Wechat.Service.Lock.RUnlock()
	message, err := PostRequest(WECHAT_SERVER_URL+"/appchat/send", map[string]interface{}{
		"access_token": config.ConfigData.Wechat.Service.AccessToken,
	}, map[string]interface{}{
		"chatid":  groupID,
		"msgtype": "markdown",
		"markdown": map[string]interface{}{
			"content": message,
		},
		"safe": 0,
	})
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(message), &data); err != nil {
		return nil
	}
	if val, ok := data["errcode"]; ok {
		switch val.(float64) {
		case 0:
			return nil
		case 40014:
			if err := GetAccessToken(); err == nil {
				if err := SendGroupMessage(groupID, message); err == nil {
					return nil
				}
			}
			return err
		default:
			return err
		}
	} else {
		return err
	}
}
