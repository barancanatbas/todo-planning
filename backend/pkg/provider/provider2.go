package provider

import (
	"encoding/json"
	"task-manager/internal/model/entity"
)

type Provider2Task struct {
	Value             int    `json:"value"`
	EstimatedDuration int    `json:"estimated_duration"`
	ID                string `json:"id"`
}

func (m *Provider2Task) ConvertAll(body []byte) ([]entity.Task, error) {
	var provider2Tasks []Provider2Task
	err := json.Unmarshal(body, &provider2Tasks)
	if err != nil {
		return nil, err
	}

	var tasks = make([]entity.Task, 0, len(provider2Tasks))
	for _, task := range provider2Tasks {
		tasks = append(tasks, entity.Task{Difficulty: task.Value, EstimatedDuration: task.EstimatedDuration, Weight: task.Value * task.EstimatedDuration, Name: task.ID})
	}
	return tasks, nil
}
