package dbHelper

import (
	"fmt"
	"time"

	"github.com/prakhar0009/go-todo/database"
	"github.com/prakhar0009/go-todo/models"
)

func IsUserExist(username string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS (SELECT * FROM users WHERE username = TRIM(LOWER($1)) AND archived_at IS NULL)`
	err := database.Todo.Get(&exists, query, username)
	return exists, err
}

func CreateUser(email, username, password string) error {
	_, err := database.Todo.Exec(
		"INSERT INTO users (email, username, password) VALUES ($1, $2, $3)",
		email, username, password,
	)
	return err
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	err := database.Todo.Get(
		&user,
		"SELECT id, email, username, password FROM users WHERE email=$1 AND archived_at IS NULL",
		email,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func CreateUserSession(userID string, expiry time.Time) (string, error) {
	var sessionID string

	err := database.Todo.QueryRow(
		"INSERT INTO user_session (user_id, expires_at) VALUES ($1, $2) RETURNING id",
		userID, expiry,
	).Scan(&sessionID)

	return sessionID, err
}

func DeleteUserSession(sessionID string) error {
	res, err := database.Todo.Exec(
		"UPDATE user_session SET archived_at = NOW() WHERE id=$1 AND archived_at IS NULL",
		sessionID,
	)

	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	fmt.Println("ROWS UPDATED:", rows)

	return nil
}
