package dbHelper

import (
	"fmt"
	"time"

	"github.com/prakhar0009/go-todo/database"
	"github.com/prakhar0009/go-todo/models"
)

func CreateTodo(userID, title, description string, expiresAt time.Time) (string, error) {
	var todoID string
	query := `
				INSERT INTO todo(
				                 user_id,
				                 title,
				                 description,
				                 expires_at,
				                 is_incomplete,
				                 is_pending
								 )
				VALUES ($1, $2, $3, $4, true, false) RETURNING id`
	err := database.Todo.Get(&todoID, query, userID, title, description, expiresAt)
	if err != nil {
		fmt.Println("DATABASE ERROR:", err)
		return "", err
	}
	return todoID, nil
}

func GetTodos(userID string, status string, limit int, offset int) ([]models.Todo, error) {
	todos := make([]models.Todo, 0)

	query := `
				SELECT
			      	id,
			      	user_id,
			      	title,
			      	description,
			      	is_completed,
			      	is_incomplete,
			      	is_pending,
			      	expires_at,
			      	created_at 
	          	FROM todo 
	          	WHERE user_id = $1 AND archived_at IS NULL`

	var args []interface{}
	args = append(args, userID)

	if status != "" {
		switch status {
		case "completed":
			query += " AND is_completed = true"
		case "incomplete":
			query += " AND is_completed = false AND expires_at > NOW()"
		case "pending":
			query += " AND is_completed = false AND expires_at <= NOW()"
		default:
			query += " AND 1=0"
		}
	}

	//query += " ORDER BY created_at DESC"
	query += fmt.Sprintf(" ORDER BY created_at DESC LIMIT %d OFFSET %d", limit, offset)

	err := database.Todo.Select(&todos, query, args...)
	return todos, err
}

func GetTodoBYID(userID, id string) (models.Todo, error) {
	var todo models.Todo
	query := `
				SELECT 
				    id,
				    user_id,
				    title,
				    description,
				    is_completed,
				    is_incomplete,
				    is_pending,
				    expires_at,
				    created_at 
	          	FROM todo
	          	WHERE user_id = $1 AND id = $2 AND archived_at IS NULL
	          	`
	err := database.Todo.Get(&todo, query, userID, id)
	return todo, err
}

func UpdateTodo(todo models.Todo) (models.Todo, error) {
	var updated models.Todo
	query := `
				UPDATE todo 
				SET title = COALESCE(NULLIF($1, ''), title),
					description = COALESCE(NULLIF($2, ''), description),
					is_completed = $3, 
					is_incomplete = $4, 
					is_pending = $5
				WHERE id = $6 AND user_id = $7 AND archived_at IS NULL
				RETURNING id, user_id, title, description, is_completed, is_incomplete, is_pending, expires_at, created_at`

	err := database.Todo.Get(&updated, query,
		todo.Title, todo.Description, todo.IsCompleted,
		todo.IsIncomplete, todo.IsPending, todo.ID, todo.UserID)
	return updated, err
}

func DeleteTodo(userID, id string) error {
	query := `
				UPDATE todo
				SET archived_at = NOW()
				WHERE id = $1 AND user_id = $2
				`
	_, err := database.Todo.Exec(query, id, userID)
	return err
}
