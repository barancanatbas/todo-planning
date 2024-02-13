package dto

import "task-manager/internal/model/entity"

type Task struct {
	ID                uint   `json:"id"`
	Name              string `json:"name"`
	Difficulty        int    `json:"difficulty"`
	EstimatedDuration int    `json:"estimatedDuration"`
}

func (d *Task) Mapping(task entity.Task) {
	d.ID = task.ID
	d.EstimatedDuration = task.EstimatedDuration
	d.Name = task.Name
	d.Difficulty = task.Difficulty
}
