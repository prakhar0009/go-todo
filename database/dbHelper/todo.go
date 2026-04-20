package dbHelper

import (
	"time"

	"github.com/prakhar0009/go-todo/database"
	"github.com/prakhar0009/go-todo/models"
)

func CreateTodo(userID, title, description string, expiresAt time.Time) (string, error) {
	var todoID string
	query := `INSERT INTO todo(user_id, title, description, expires_at) VALUES ($1, $2, $3, $4) RETURNING id`
	err := database.Todo.Get(&todoID, query, userID, title, description, expiresAt)
	if err != nil {
		return "", err
	}
	return todoID, nil
}

func GetAllTodo(userID string) (models.Todo, error) {
	var todo models.Todo
	query := `SELECT * FROM todo WHERE user_id = $1 AND archived_at IS NULL`
	err := database.Todo.Get(&todo, query, userID)
	return todo, err
}

func GetTodoBYID(userID, id string) (models.Todo, error) {
	var todo models.Todo
	query := `SELECT * FROM todo WHERE user_id = $1 AND id = $2 AND archive_at IS NULL`
	err := database.Todo.Get(&todo, query, userID, id)
	return todo, err
}

func UpdateTodo(userID, id, title, description string) (models.Todo, error) {
	var updatedTodo models.Todo
	query := `UPDATE todo SET title = $1, description = $2 WHERE id = $3 AND user_id = $4 AND archived_at IS NULL RETURNING *`
	_, err := database.Todo.Exec(query, title, description, id, userID)
	return updatedTodo, err
}

func DeleteTodo(userID, id string) error {
	query := `UPDATE todo SET archive_at = NOW() WHERE id = $1 AND user_id = $2`
	_, err := database.Todo.Exec(query, id, userID)
	return err
}
