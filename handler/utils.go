package handler

import "github.com/Zheaoli/nexus/config"

func GetGroupID(sourceurl string) string {
	return config.URLMap[sourceurl].Wechat.GroupID
}
