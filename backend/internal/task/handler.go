package task

import (
	"github.com/gofiber/fiber/v2"
	"task-manager/internal/model/dto"
)

type handler struct {
	service IService
}

func NewTaskHandler(service IService) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) GetPlanning(c *fiber.Ctx) error {
	developers, err := h.service.GetPlanning()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.FailedResponse{
			Message: "An unexpected error occurred",
			Error:   err.Error(),
		})
	}

	numberOfWeek, err := h.service.GetNumberOfWeek()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.FailedResponse{
			Message: "An unexpected error occurred",
			Error:   err.Error(),
		})
	}

	var response dto.PlanningResponse
	response.Mapping(developers)
	response.TotalWeek = numberOfWeek.NumberOfWeeks
	return c.Status(fiber.StatusOK).JSON(response)
}
