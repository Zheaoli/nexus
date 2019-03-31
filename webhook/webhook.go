package webhook

type WBStrategy interface {
	ParseMessage() (WBMessage, error)
}

type WBMessage interface {
	GetMessageType() string
	GetMessageContent()
}
