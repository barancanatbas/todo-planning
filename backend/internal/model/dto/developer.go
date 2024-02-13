package dto

import "task-manager/internal/model/entity"

type Developer struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Tasks []Task `json:"tasks"`
}

func (d *Developer) Mapping(developer entity.Developer) {
	d.ID = developer.ID
	d.Name = developer.Name

	for _, task := range developer.Tasks {
		responseTask := Task{}
		responseTask.Mapping(task)
		d.Tasks = append(d.Tasks, responseTask)
	}
}
