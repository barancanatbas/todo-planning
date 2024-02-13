package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"net/url"
	"task-manager/internal/devloper"
	"task-manager/internal/task"
	"task-manager/pkg/provider"
)

type router struct {
	app *fiber.App
}

func NewRouter(app *fiber.App) *router {
	return &router{
		app: app,
	}
}

func (r *router) RegisterRoutes(db *gorm.DB) {
	var taskHandler ITaskHandler

	providerClient := provider.NewProviderClient()

	developerRepository := devloper.NewDeveloperRepository(db)
	developerService := devloper.NewDeveloperService(developerRepository)

	taskRepository := task.NewTaskRepository(db)
	taskService := task.NewTaskService(r.getProviderMap(), developerService, providerClient, taskRepository)
	taskHandler = task.NewTaskHandler(taskService)

	r.taskHandler(taskHandler)
}

func (r *router) getProviderMap() map[*url.URL]provider.ITask {
	providerMap := make(map[*url.URL]provider.ITask)

	provider1Url, err := url.Parse("https://run.mocky.io/v3/27b47d79-f382-4dee-b4fe-a0976ceda9cd")
	if err != nil {
		log.Fatal().Err(err)
	}
	providerMap[provider1Url] = &provider.Provider1Task{}

	provider2Url, err := url.Parse("https://run.mocky.io/v3/7b0ff222-7a9c-4c54-9396-0df58e289143")
	if err != nil {
		log.Fatal().Err(err)
	}
	providerMap[provider2Url] = &provider.Provider2Task{}
	return providerMap
}

func (r *router) taskHandler(taskHandler ITaskHandler) {
	taskGroup := r.app.Group("/tasks")

	taskGroup.Get("/", taskHandler.GetPlanning)
}
