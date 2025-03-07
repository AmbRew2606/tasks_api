package routes

import (
	"tz_todo_list_1/internal/handlers"
	"tz_todo_list_1/pkg/storage/postgres"

	"github.com/gofiber/fiber/v2"
)

func TaskRoutes(app *fiber.App, storage *postgres.Storage) {
	handler := handlers.NewTaskHandler(storage)

	// Создание задачи
	app.Post("/tasks", handler.CreateTask)

	// Вывод всех задач
	app.Get("/tasks", handler.GetTasks)

	// Обновление задачи
	app.Put("/tasks/:id", handler.UpdateTask)

	// Удаление задачи
	app.Delete("/tasks/:id", handler.DeleteTask)
}
