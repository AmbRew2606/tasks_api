package postgres

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"tz_todo_list_1/config"
	"tz_todo_list_1/internal/models"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Пул соединений с БД psql
type Storage struct {
	db *pgxpool.Pool
}

// New - создает подключение к БД
func New() (*Storage, error) {

	cfg := config.LoadConfig()

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)

	dbpool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("postgres: ошибка подключения к БД: %w", err)
	}

	return &Storage{db: dbpool}, nil
}

// Close - закрытие подключения с БД
func (s *Storage) Close() {
	s.db.Close()
}

// GetTasks - вывоод всех задач
func (s *Storage) GetTasks() ([]models.Task, error) {
	var tasks []models.Task
	rows, err := s.db.Query(context.Background(), "SELECT id, title, description, status, created_at, updated_at FROM tasks")
	if err != nil {
		return nil, fmt.Errorf("postgres: ошибка выполнения запроса: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt); err != nil {
			return nil, fmt.Errorf("postgres: ошибка при сканировании строки: %w", err)
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("postgres: ошибка при обработке строк: %w", err)
	}

	return tasks, nil
}

// CreateTask - создание новой задачи
func (s *Storage) CreateTask(t models.Task) (int, error) {
	var taskID int
	err := s.db.QueryRow(context.Background(), `
		INSERT INTO tasks (title, description, status) 
		VALUES ($1, $2, $3) RETURNING id;
	`, t.Title, t.Description, t.Status).Scan(&taskID)

	if err != nil {
		return 0, fmt.Errorf("postgres: ошибка при создании задачи: %w", err)
	}
	return taskID, nil
}

// UpdateTask - обновляет задачу
// Логика метода позволяет обновлять любое поле
func (s *Storage) UpdateTask(t models.Task) error {

	query := `UPDATE tasks SET updated_at = now()`
	var args []interface{}
	var setClauses []string

	if t.Title != "" {
		setClauses = append(setClauses, "title = $"+strconv.Itoa(len(args)+1))
		args = append(args, t.Title)
	}
	if t.Description != "" {
		setClauses = append(setClauses, "description = $"+strconv.Itoa(len(args)+1))
		args = append(args, t.Description)
	}
	if t.Status != "" {
		setClauses = append(setClauses, "status = $"+strconv.Itoa(len(args)+1))
		args = append(args, t.Status)
	}

	if len(setClauses) == 0 {
		return fmt.Errorf("postgres: нет данных для обновления")
	}

	// "сбор" запроса
	query += ", " + strings.Join(setClauses, ", ") + " WHERE id = $" + strconv.Itoa(len(args)+1)
	args = append(args, t.ID)

	_, err := s.db.Exec(context.Background(), query, args...)
	if err != nil {
		return fmt.Errorf("postgres: ошибка при обновлении задачи: %w", err)
	}
	return nil
}

// DeleteTask - удаление задачи по ID
func (s *Storage) DeleteTask(taskID int) error {
	_, err := s.db.Exec(context.Background(), `
		DELETE FROM tasks WHERE id = $1;
	`, taskID)

	if err != nil {
		return fmt.Errorf("postgres: ошибка при удалении задачи: %w", err)
	}
	return nil
}
