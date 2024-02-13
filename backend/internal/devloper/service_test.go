package devloper

import (
	"github.com/stretchr/testify/assert"
	"task-manager/internal/model/entity"
	mocks "task-manager/mocks/internal_/devloper"
	"testing"
)

func TestGetDevelopers(t *testing.T) {
	mockRepo := mocks.NewIRepository(t)

	mockRepo.On("FindAll").Return([]entity.Developer{}, nil)

	devService := NewDeveloperService(mockRepo)
	developers, err := devService.GetDevelopers()

	assert.NoError(t, err)
	assert.NotNil(t, developers)

	mockRepo.AssertExpectations(t)
}

func TestGetDevelopersAndTasks(t *testing.T) {
	mockRepo := mocks.NewIRepository(t)
	mockRepo.On("FindAllAndTask").Return([]entity.Developer{}, nil)

	devService := NewDeveloperService(mockRepo)
	developers, err := devService.GetDevelopersAndTasks()

	assert.NoError(t, err)
	assert.NotNil(t, developers)

	mockRepo.AssertExpectations(t)
}
