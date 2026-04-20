package dbHelper

import (
	"time"

	"github.com/prakhar0009/go-todo/database"
	"github.com/prakhar0009/go-todo/models"
)

func CreateTodo(userID, title, description string, expiresAt time.Time) (string, error) {
	var todoID string
	query := `
				INSERT INTO todo(user_id, title, description, expires_at, is_incomplete)
				VALUES ($1, $2, $3, $4, true) RETURNING id`
	err := database.Todo.Get(&todoID, query, userID, title, description, expiresAt)
	if err != nil {
		return "", err
	}
	return todoID, nil
}

func GetAllTodo(userID string) ([]models.Todo, error) {
	var todos []models.Todo
	query := `
				SELECT id, user_id, title, description, is_completed, is_incomplete, is_pending, expires_at, created_at 
	          	FROM todo WHERE user_id = $1 AND id = $2 AND archived_at IS NULL
	          	`
	err := database.Todo.Select(&todos, query, userID)
	return todos, err
}

func GetTodoBYID(userID, id string) (models.Todo, error) {
	var todo models.Todo
	query := `
				SELECT id, user_id, title, description, is_completed, is_incomplete, is_pending, expires_at, created_at 
	          	FROM todo WHERE user_id = $1 AND id = $2 AND archived_at IS NULL
	          	`
	err := database.Todo.Get(&todo, query, userID, id)
	return todo, err
}

func UpdateTodo(todo models.Todo) (models.Todo, error) {
	var updated models.Todo
	query := `
				UPDATE todo 
				SET title = $1, description = $2, is_completed = $3, is_incomplete = $4, is_pending = $5
				WHERE id = $6 AND user_id = $7 AND archived_at IS NULL
				RETURNING id, user_id, title, description, is_completed, is_incomplete, is_pending, expires_at, created_at
				`
	err := database.Todo.Get(&updated, query,
		todo.Title, todo.Description, todo.IsCompleted,
		todo.IsIncomplete, todo.IsPending, todo.ID, todo.UserID)
	return updated, err
}

func DeleteTodo(userID, id string) error {
	query := `
				UPDATE todo
				SET archive_at = NOW()
				WHERE id = $1 AND user_id = $2
				`
	_, err := database.Todo.Exec(query, id, userID)
	return err
}
