package tower

type TWMessageEvent interface {
	GetEventType() string
	Parse() (string, error)
}

type TWMessage struct {
	Event TWMessageEvent
}

func (tw *TWMessage) String() string {
	return "Tower message"
}

func (tw *TWMessage) Parse() string {

}
