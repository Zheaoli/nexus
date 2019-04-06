package event

import (
	"encoding/json"
	"fmt"
	"github.com/Zheaoli/nexus/config"
	"io/ioutil"
	"path"
)

type TodoListsEvent struct {
	TemplatePath string
	EventData    TodoListEventBody
}
type TodoListEventBody struct {
	Action string            `json:"action"`
	Data   TodoListEventData `json:"data"`
}
type TodoListEventData struct {
	Project  ProjectEvent         `json:"project"`
	TodoList TodoListEventContent `json:"todolist"`
}
type TodoListEventContent struct {
	Guid      string               `json:"guid"`
	Name      string               `json:"name"`
	UpdatedAt string               `json:"updated_at"`
	Handler   TodoListEventHandler `json:"handler"`
}
type TodoListEventHandler struct {
	Guid     string `json:"guid"`
	NickName string `json:"nickname"`
}

func New(message string, url string) (*TodoListsEvent, error) {
	var todoList TodoListEventBody
	if err := json.Unmarshal([]byte(message), &todoList); err != nil {
		return nil, err
	}
	templatePath := ""
	if val, ok := config.URLMap[url]; ok {
		templatePath = val.(*config.WebhookConfig).TemplatePath
	} else {
		return nil, ConfigNotExist
	}

	return &TodoListsEvent{TemplatePath: templatePath, EventData: todoList}, nil
}

func (te *TodoListsEvent) GetEventType() string {
	return "todolists"
}

func (te *TodoListsEvent) Parse() (string, error) {
	filePath := path.Join(te.TemplatePath, "todolists", te.EventData.Action, "index.md")
	templateData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	result := fmt.Sprintf(string(templateData), te.EventData.Data.Project.Name,
		te.EventData.Data.TodoList.Handler.NickName, te.EventData.Data.TodoList.Name)
	return result, nil
}
