package handler

import (
	"github.com/Zheaoli/nexus/requests/wechat"
	"github.com/Zheaoli/nexus/webhook"
	"github.com/Zheaoli/nexus/webhook/tower"
	"io"
	"io/ioutil"
	"net/http"
)

type UrlParser struct {
	Source string `url:"source" binding:"required"`
	Target string `url:"target" binding:"required"`
	URL    string `url:"url" binding:"required"`
}

func HandleWebHook(source, target, sourceurl string, body io.ReadCloser, headers http.Header) error {
	var sourceHandler webhook.WBStrategy
	message, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	switch source {
	case "tower":
		sourceHandler = tower.New(string(message), headers.Get("X-Tower-Event"), sourceurl)
	default:
		return nil
	}
	messageHandler, err := sourceHandler.ParseMessage()
	if err != nil {
		return err
	}
	messageData, err := messageHandler.Parse()
	if err != nil {
		return err
	}
	switch target {
	case "wechat":
		err = wechat.SendGroupMessage(GetGroupID(sourceurl), messageData)
		if err != nil {
			return err
		}
		return nil
	default:
		return nil
	}

}
