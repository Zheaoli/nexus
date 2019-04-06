package tower_event

import (
	"encoding/json"
	"fmt"
	"github.com/Zheaoli/nexus/config"
	"io/ioutil"
	"path"
)

type TodoEvent struct {
	TemplatePath string
	EventData    TodoEventBody
}
type TodoEventBody struct {
	Action string        `json:"action"`
	Data   TodoEventData `json:"data"`
}
type TodoEventData struct {
	Project ProjectEvent     `json:"project"`
	Todo    TodoEventContent `json:"todo"`
}
type TodoEventContent struct {
	Guid      string            `json:"guid"`
	Title     string            `json:"title"`
	UpdatedAt string            `json:"updated_at"`
	Handler   TodoEventHandler  `json:"handler"`
	Assignee  TodoEventAssignee `json:"assignee"`
	DueAt     string            `json:"due_at"`
}
type TodoEventHandler struct {
	Guid     string `json:"guid"`
	NickName string `json:"nickname"`
}
type TodoEventAssignee struct {
	Guid     string `json:"guid"`
	NickName string `json:"nickname"`
}

const (
	TODOMoved    = "moved"
	TODOAssign   = "assigned"
	TODODeadLine = "deadline_changed"
)

func NewTodo(message string, url string) (*TodoEvent, error) {
	var todoList TodoEventBody
	if err := json.Unmarshal([]byte(message), &todoList); err != nil {
		return nil, err
	}
	templatePath := ""
	if val, ok := config.URLMap[url]; ok {
		templatePath = val.TemplatePath
	} else {
		return nil, ConfigNotExist
	}

	return &TodoEvent{TemplatePath: templatePath, EventData: todoList}, nil
}

func (te *TodoEvent) GetEventType() string {
	return "todolists"
}

func (te *TodoEvent) Parse() (string, error) {
	filePath := path.Join(te.TemplatePath, "todos", te.EventData.Action, "index.md")
	templateData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	var result string
	switch te.EventData.Action {
	case TODOMoved:
		result = "abc"
	case TODOAssign:
		result = fmt.Sprintf(string(templateData), te.EventData.Data.Project.Name,
			te.EventData.Data.Todo.Handler.NickName, te.EventData.Data.Todo.Title,
			te.EventData.Data.Todo.Assignee.NickName)
	case TODODeadLine:
		result=fmt.Sprintf(string(templateData), te.EventData.Data.Project.Name,
			te.EventData.Data.Todo.Handler.NickName, te.EventData.Data.Todo.Title,
			te.EventData.Data.Todo.DueAt)
	default:
		result = fmt.Sprintf(string(templateData), te.EventData.Data.Project.Name,
			te.EventData.Data.Todo.Handler.NickName, te.EventData.Data.Todo.Title)
	}

	return result, nil
}
