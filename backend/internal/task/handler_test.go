package task

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"task-manager/internal/model/entity"
	mocks "task-manager/mocks/internal_/task"
	"testing"
)

func TestHandler_GetPlanning(t *testing.T) {
	mockService := mocks.NewIService(t)

	handler := NewTaskHandler(mockService)

	mockService.On("GetPlanning").Return([]entity.Developer{}, nil)

	mockService.On("GetNumberOfWeek").Return(entity.Week{}, nil)

	app := fiber.New()

	app.Get("/planning", handler.GetPlanning)

	req, _ := http.NewRequest("GET", "/planning", nil)

	resp, _ := app.Test(req)

	mockService.AssertExpectations(t)

	assert.Equal(t, 200, resp.StatusCode)
}

func TestHandler_GetPlanning_ShouldReturnGetPlanningError(t *testing.T) {
	mockService := mocks.NewIService(t)

	handler := NewTaskHandler(mockService)

	mockService.On("GetPlanning").Return(nil, errors.New("internal server error"))

	app := fiber.New()

	app.Get("/planning", handler.GetPlanning)

	req, _ := http.NewRequest("GET", "/planning", nil)

	resp, _ := app.Test(req)

	mockService.AssertExpectations(t)

	assert.Equal(t, 400, resp.StatusCode)
}

func TestHandler_GetPlanning_ShouldReturnGetNumberOfWeekError(t *testing.T) {
	mockService := mocks.NewIService(t)

	handler := NewTaskHandler(mockService)

	mockService.On("GetPlanning").Return([]entity.Developer{}, nil)

	mockService.On("GetNumberOfWeek").Return(entity.Week{}, errors.New("not found error"))
	app := fiber.New()

	app.Get("/planning", handler.GetPlanning)

	req, _ := http.NewRequest("GET", "/planning", nil)

	resp, _ := app.Test(req)

	mockService.AssertExpectations(t)

	assert.Equal(t, 400, resp.StatusCode)
}
