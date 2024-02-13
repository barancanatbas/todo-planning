package provider

import (
	"encoding/json"
	"task-manager/internal/model/entity"
)

type Provider1Task struct {
	Zorluk int    `json:"zorluk"`
	Sure   int    `json:"sure"`
	ID     string `json:"id"`
}

func (m *Provider1Task) ConvertAll(body []byte) ([]entity.Task, error) {
	var provider1Tasks []Provider1Task
	err := json.Unmarshal(body, &provider1Tasks)
	if err != nil {
		return nil, err
	}

	var tasks = make([]entity.Task, 0, len(provider1Tasks))
	for _, task := range provider1Tasks {
		tasks = append(tasks, entity.Task{Difficulty: task.Zorluk, EstimatedDuration: task.Sure, Weight: task.Zorluk * task.Sure, Name: task.ID})
	}
	return tasks, nil
}
