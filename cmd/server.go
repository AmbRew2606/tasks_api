package main

import (
	"log"
	"os"

	"tz_todo_list_1/internal/routes"
	"tz_todo_list_1/pkg/storage/postgres"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	// Загрузка переменных ищ  env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка при загрузке .env файла")
	}

	// Подключение к базе данных
	storage, err := postgres.New()
	if err != nil {
		log.Fatal("Ошибка при подключении к базе данных:", err)
	}
	defer storage.Close()

	app := fiber.New()

	// Регистрация маршрутов
	routes.TaskRoutes(app, storage)

	// Запуск сервера
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	err = app.Listen(":" + port)
	if err != nil {
		log.Fatal("Ошибка при запуске сервера:", err)
	}
}
