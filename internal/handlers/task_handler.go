package handlers

import (
	"log"
	"strconv"
	"tz_todo_list_1/internal/models"
	"tz_todo_list_1/pkg/storage/postgres"

	"github.com/gofiber/fiber/v2"
)

type TaskHandler struct {
	Storage *postgres.Storage
}

func NewTaskHandler(storage *postgres.Storage) *TaskHandler {
	return &TaskHandler{Storage: storage}
}

// CreateTask (POST /tasks).
func (h *TaskHandler) CreateTask(c *fiber.Ctx) error {
	var task models.Task

	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Недопустимый текст запроса",
		})
	}

	taskID, err := h.Storage.CreateTask(task)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не удалось создать задачу",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Задача успешно создана",
		"task_id": taskID,
	})
}

// GetTasks (GET /tasks).
func (h *TaskHandler) GetTasks(c *fiber.Ctx) error {

	tasks, err := h.Storage.GetTasks()
	if err != nil {
		log.Printf("Ошибка при получении задач: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не удалось получить список задач",
		})
	}

	return c.JSON(tasks)
}

// UpdateTask (PUT /tasks/:id).
func (h *TaskHandler) UpdateTask(c *fiber.Ctx) error {
	taskID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Неверный ID задачи",
		})
	}

	var task models.Task

	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Недопустимый текст запроса",
		})
	}

	task.ID = taskID

	if err := h.Storage.UpdateTask(task); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не удалось обновить задачу",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Задача успешно обновлена",
	})
}

// DeleteTask (DELETE /tasks/:id).
func (h *TaskHandler) DeleteTask(c *fiber.Ctx) error {
	taskID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Неверный ID задачи",
		})
	}

	if err := h.Storage.DeleteTask(taskID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не удалось удалить задачу",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Задача успешно удалена",
	})
}
