package task

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/url"
	"task-manager/internal/model/entity"
	mocks "task-manager/mocks/internal_/task"
	"task-manager/pkg/provider"
	"testing"
)

func TestService_FetchAll_ShouldReturnSuccess(t *testing.T) {
	providerMap := make(map[*url.URL]provider.ITask)

	provider1Url, err := url.Parse("https://run.mocky.io/v3/27b47d79-f382-4dee-b4fe-a0976ceda9cd")
	if err != nil {
		panic(err)
	}
	providerMap[provider1Url] = &provider.Provider1Task{}

	mockRepository := mocks.NewITaskRepository(t)
	mockProviderClient := mocks.NewIProviderClient(t)
	mockDeveloperService := mocks.NewIDeveloperService(t)

	tasks := []entity.Task{
		{Name: "task 1", EstimatedDuration: 10, Difficulty: 3, Weight: 30},
		{Name: "task 2", EstimatedDuration: 10, Difficulty: 2, Weight: 20},
		{Name: "task 3", EstimatedDuration: 10, Difficulty: 5, Weight: 50},
	}

	developers := []entity.Developer{
		{Name: "DEVELOPER 1", Capacity: 45, Tasks: []entity.Task{}},
		{Name: "DEVELOPER 2", Capacity: 90, Tasks: []entity.Task{}},
		{Name: "DEVELOPER 3", Capacity: 135, Tasks: []entity.Task{}},
		{Name: "DEVELOPER 4", Capacity: 180, Tasks: []entity.Task{}},
		{Name: "DEVELOPER 5", Capacity: 225, Tasks: []entity.Task{}},
	}

	mockProviderClient.On("Get", provider1Url, providerMap[provider1Url]).Return(tasks, nil)
	mockRepository.On("Save", mock.Anything).Return(nil)
	mockRepository.Mock.On("SaveNumberOfWeek", 1).Return(nil)
	mockDeveloperService.On("GetDevelopers").Return(developers, nil)

	service := NewTaskService(providerMap, mockDeveloperService, mockProviderClient, mockRepository)

	err = service.FetchAll()

	assert.NoError(t, err)
}

func TestService_FetchAll_WithWrongProvider_ShouldReturnError(t *testing.T) {
	providerMap := make(map[*url.URL]provider.ITask)

	provider1Url, err := url.Parse("https://run.mocky.io/v3/27b47d79-f382-4dee-b4fe-a0976ceda9cd")
	if err != nil {
		panic(err)
	}
	providerMap[provider1Url] = &provider.Provider1Task{}

	mockProviderClient := mocks.NewIProviderClient(t)

	mockProviderClient.On("Get", provider1Url, providerMap[provider1Url]).Return(nil, errors.New("provider error"))

	service := NewTaskService(providerMap, nil, mockProviderClient, nil)

	err = service.FetchAll()

	assert.Error(t, err)
}

func TestService_FetchAll_ShouldDeveloperError(t *testing.T) {
	providerMap := make(map[*url.URL]provider.ITask)

	provider1Url, err := url.Parse("https://run.mocky.io/v3/27b47d79-f382-4dee-b4fe-a0976ceda9cd")
	if err != nil {
		panic(err)
	}
	providerMap[provider1Url] = &provider.Provider1Task{}

	mockProviderClient := mocks.NewIProviderClient(t)
	mockDeveloperService := mocks.NewIDeveloperService(t)

	tasks := []entity.Task{
		{Name: "task 1", EstimatedDuration: 10, Difficulty: 3, Weight: 30},
		{Name: "task 2", EstimatedDuration: 10, Difficulty: 2, Weight: 20},
		{Name: "task 3", EstimatedDuration: 10, Difficulty: 5, Weight: 50},
	}

	mockProviderClient.On("Get", provider1Url, providerMap[provider1Url]).Return(tasks, nil)
	mockDeveloperService.On("GetDevelopers").Return(nil, errors.New("developers not found"))

	service := NewTaskService(providerMap, mockDeveloperService, mockProviderClient, nil)

	err = service.FetchAll()

	assert.Error(t, err)
}

func TestService_FetchAll_ShouldTaskSaveError(t *testing.T) {
	providerMap := make(map[*url.URL]provider.ITask)

	provider1Url, err := url.Parse("https://run.mocky.io/v3/27b47d79-f382-4dee-b4fe-a0976ceda9cd")
	if err != nil {
		panic(err)
	}
	providerMap[provider1Url] = &provider.Provider1Task{}

	mockRepository := mocks.NewITaskRepository(t)
	mockProviderClient := mocks.NewIProviderClient(t)
	mockDeveloperService := mocks.NewIDeveloperService(t)

	tasks := []entity.Task{
		{Name: "task 1", EstimatedDuration: 10, Difficulty: 3, Weight: 30},
		{Name: "task 2", EstimatedDuration: 10, Difficulty: 2, Weight: 20},
		{Name: "task 3", EstimatedDuration: 10, Difficulty: 5, Weight: 50},
	}

	developers := []entity.Developer{
		{Name: "DEVELOPER 1", Capacity: 45, Tasks: []entity.Task{}},
		{Name: "DEVELOPER 2", Capacity: 90, Tasks: []entity.Task{}},
		{Name: "DEVELOPER 3", Capacity: 135, Tasks: []entity.Task{}},
		{Name: "DEVELOPER 4", Capacity: 180, Tasks: []entity.Task{}},
		{Name: "DEVELOPER 5", Capacity: 225, Tasks: []entity.Task{}},
	}

	mockProviderClient.On("Get", provider1Url, providerMap[provider1Url]).Return(tasks, nil)
	mockRepository.On("Save", mock.Anything).Return(errors.New("save error"))
	mockDeveloperService.On("GetDevelopers").Return(developers, nil)

	service := NewTaskService(providerMap, mockDeveloperService, mockProviderClient, mockRepository)

	err = service.FetchAll()

	assert.Error(t, err)
}

func TestService_GetPlanning_ShouldSuccess(t *testing.T) {
	expectedDevelopers := []entity.Developer{
		{Name: "DEVELOPER 1", Capacity: 45, Tasks: []entity.Task{}},
		{Name: "DEVELOPER 2", Capacity: 90, Tasks: []entity.Task{}},
		{Name: "DEVELOPER 3", Capacity: 135, Tasks: []entity.Task{}},
		{Name: "DEVELOPER 4", Capacity: 180, Tasks: []entity.Task{}},
		{Name: "DEVELOPER 5", Capacity: 225, Tasks: []entity.Task{}},
	}

	mockDeveloperService := mocks.NewIDeveloperService(t)
	mockDeveloperService.On("GetDevelopersAndTasks").Return(expectedDevelopers, nil)

	service := &service{
		developerService: mockDeveloperService,
	}

	result, err := service.GetPlanning()

	assert.NoError(t, err)
	assert.Equal(t, result, expectedDevelopers)
}
