package devloper

import "task-manager/internal/model/entity"

type IRepository interface {
	FindAll() ([]entity.Developer, error)
	FindAllAndTask() ([]entity.Developer, error)
}
