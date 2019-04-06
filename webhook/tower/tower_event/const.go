package tower_event

import "errors"

type ProjectEvent struct {
	Guid string `json:"guid"`
	Name string `json:"name"`
}

var ConfigNotExist = errors.New("config not exist")

const (
	TodoListsEventF = "todolists"
	TodosEventF      = "todos"
)
