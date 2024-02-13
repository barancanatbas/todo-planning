package router

import "github.com/gofiber/fiber/v2"

type ITaskHandler interface {
	GetPlanning(c *fiber.Ctx) error
}
