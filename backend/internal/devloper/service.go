package devloper

import "task-manager/internal/model/entity"

type service struct {
	repository IRepository
}

func NewDeveloperService(repository IRepository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetDevelopers() ([]entity.Developer, error) {
	return s.repository.FindAll()
}

func (s *service) GetDevelopersAndTasks() ([]entity.Developer, error) {
	return s.repository.FindAllAndTask()
}
