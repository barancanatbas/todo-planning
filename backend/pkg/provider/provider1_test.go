package provider

import (
	"github.com/stretchr/testify/assert"
	"task-manager/internal/model/entity"
	"testing"
)

func TestProvider1Task_ConvertAll_ShouldSuccess(t *testing.T) {
	body := []byte(`[
		{"zorluk": 1, "sure": 2, "id": "task1"},
		{"zorluk": 3, "sure": 4, "id": "task2"}
	]`)

	p1t := &Provider1Task{}

	tasks, err := p1t.ConvertAll(body)

	assert.NoError(t, err)

	expectedTasks := []entity.Task{
		{Difficulty: 1, EstimatedDuration: 2, Weight: 2, Name: "task1"},
		{Difficulty: 3, EstimatedDuration: 4, Weight: 12, Name: "task2"},
	}
	assert.Equal(t, expectedTasks, tasks)
}

func TestProvider1Task_ConvertAll_ShouldReturnJsonError(t *testing.T) {
	body := []byte(`invalid json`)

	p1t := &Provider1Task{}

	_, err := p1t.ConvertAll(body)
	assert.Error(t, err)
}
