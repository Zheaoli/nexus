package tower

import "github.com/Zheaoli/nexus/webhook/tower/tower_event"

type TWMessageEvent interface {
	GetEventType() string
	Parse() (string, error)
}

type TWMessage struct {
	Event TWMessageEvent
}

func NewMessage(message, event, souceurl string) (*TWMessage, error) {
	eventData, err := NewEvent(message, event, souceurl)
	if err != nil {
		return nil, err
	}
	return &TWMessage{Event: eventData}, nil
}

func NewEvent(message, event, sourceurl string) (TWMessageEvent, error) {
	switch event {
	case tower_event.TodoListsEventF:
		return tower_event.NewTodoLists(message, sourceurl)
	case tower_event.TodosEventF:
		return tower_event.NewTodo(message, sourceurl)
	default:
		return nil, EventNotDefined
	}
}

func (tw *TWMessage) String() string {
	return "Tower message"
}

func (tw *TWMessage) Parse() (string, error) {
	return tw.Event.Parse()
}
