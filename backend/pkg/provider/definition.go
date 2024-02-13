package provider

import "task-manager/internal/model/entity"

type ITask interface {
	ConvertAll(body []byte) ([]entity.Task, error)
}
