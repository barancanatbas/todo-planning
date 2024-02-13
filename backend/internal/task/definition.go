package task

import (
	"net/url"
	"task-manager/internal/model/entity"
	"task-manager/pkg/provider"
)

type IService interface {
	GetPlanning() ([]entity.Developer, error)
	GetNumberOfWeek() (entity.Week, error)
}

type ITaskRepository interface {
	Save(tasks []entity.Task) error
	SaveNumberOfWeek(numberOfWeek int) error
	GetNumberOfWeek() (entity.Week, error)
}

type IDeveloperService interface {
	GetDevelopers() ([]entity.Developer, error)
	GetDevelopersAndTasks() ([]entity.Developer, error)
}

type IProviderClient interface {
	Get(url *url.URL, providerModel provider.ITask) ([]entity.Task, error)
}
