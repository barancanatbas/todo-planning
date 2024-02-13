package dto

import "task-manager/internal/model/entity"

type PlanningResponse struct {
	Developers []Developer `json:"developers"`
	TotalWeek  int         `json:"totalWeek"`
}

func (d *PlanningResponse) Mapping(developers []entity.Developer) {
	for _, developerEntity := range developers {
		responseDeveloper := Developer{}
		responseDeveloper.Mapping(developerEntity)
		d.Developers = append(d.Developers, responseDeveloper)
	}
}
