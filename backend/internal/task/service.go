package task

import (
	"github.com/rs/zerolog/log"
	"net/url"
	"sort"
	"sync"
	"task-manager/internal/model/entity"
	"task-manager/pkg/provider"
)

type service struct {
	providerMap      map[*url.URL]provider.ITask
	repository       ITaskRepository
	developerService IDeveloperService
	providerClient   IProviderClient
}

func NewTaskService(providerMap map[*url.URL]provider.ITask, developerService IDeveloperService, providerClient IProviderClient, repository ITaskRepository) *service {
	taskService := &service{
		providerMap:      providerMap,
		repository:       repository,
		developerService: developerService,
		providerClient:   providerClient,
	}

	taskService.FetchAll()
	return taskService
}

func (s *service) FetchAll() error {
	var allTasks []entity.Task
	var wg sync.WaitGroup
	errChan := make(chan error, len(s.providerMap))

	for providerUrl, providerModel := range s.providerMap {
		wg.Add(1)
		go func(url *url.URL, provider provider.ITask) {
			defer wg.Done()
			tasks, err := s.providerClient.Get(url, provider)
			if err != nil {
				log.Error().Err(err)
				errChan <- err
				return
			}
			allTasks = append(allTasks, tasks...)
		}(providerUrl, providerModel)
	}

	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			return err
		}
	}

	developers, err := s.developerService.GetDevelopers()
	if err != nil {
		log.Error().Err(err)
		return err
	}

	return s.assignTasks(allTasks, developers)
}

func (s *service) GetPlanning() ([]entity.Developer, error) {
	return s.developerService.GetDevelopersAndTasks()
}

func (s *service) GetNumberOfWeek() (entity.Week, error) {
	return s.repository.GetNumberOfWeek()
}

func (s *service) assignTasks(tasks []entity.Task, developers []entity.Developer) error {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Weight > tasks[j].Weight
	})

	sort.Slice(developers, func(i, j int) bool {
		return developers[i].Capacity > developers[j].Capacity
	})

	numberOfWeeks := 0

	for len(tasks) > 0 {
		for i, developer := range developers {
			weeklyCapacity := developer.Capacity
			for j := 0; j < len(tasks); j++ {
				task := tasks[j]
				if weeklyCapacity <= 0 {
					break
				}

				if task.Weight <= weeklyCapacity {
					weeklyCapacity -= task.Weight
					task.DeveloperID = developer.ID
					developers[i].Tasks = append(developers[i].Tasks, task)
					tasks = append(tasks[:j], tasks[j+1:]...)
					j--
					continue
				}
			}
		}
		numberOfWeeks++
	}

	err := s.saveTasks(developers)
	if err != nil {
		log.Error().Err(err)
		return err
	}

	return s.repository.SaveNumberOfWeek(numberOfWeeks)
}

func (s *service) saveTasks(developers []entity.Developer) error {
	var tasks []entity.Task
	for _, developer := range developers {
		tasks = append(tasks, developer.Tasks...)
	}

	return s.repository.Save(tasks)
}
