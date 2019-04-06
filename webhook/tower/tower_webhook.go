package tower

import "github.com/Zheaoli/nexus/webhook"

type TWWebHook struct {
	message   string
	event     string
	sourceurl string
}

func New(message string, event string, url string) *TWWebHook {
	return &TWWebHook{message: message, event: event, sourceurl: url}
}

func (tw *TWWebHook) String() string {
	return "Tower Webhook messgae"
}

func (tw *TWWebHook) ParseMessage() (webhook.WBMessage, error) {
	return NewMessage(tw.message, tw.event, tw.sourceurl)
}
