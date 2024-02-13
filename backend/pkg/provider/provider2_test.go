package provider

import (
	"github.com/stretchr/testify/assert"
	"task-manager/internal/model/entity"
	"testing"
)

func TestProvider2Task_ConvertAll_ShouldSuccess(t *testing.T) {
	body := []byte(`[
		{"value": 1, "estimated_duration": 2, "id": "task1"},
		{"value": 3, "estimated_duration": 4, "id": "task2"}
	]`)

	p2t := &Provider2Task{}

	tasks, err := p2t.ConvertAll(body)

	assert.NoError(t, err)

	expectedTasks := []entity.Task{
		{Difficulty: 1, EstimatedDuration: 2, Weight: 2, Name: "task1"},
		{Difficulty: 3, EstimatedDuration: 4, Weight: 12, Name: "task2"},
	}
	assert.Equal(t, expectedTasks, tasks)
}

func TestProvider2Task_ConvertAll_ShouldReturnJsonError(t *testing.T) {
	body := []byte(`invalid json`)

	p2t := &Provider2Task{}

	_, err := p2t.ConvertAll(body)
	assert.Error(t, err)
}
