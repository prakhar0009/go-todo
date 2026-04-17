package dbHelper

import (
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
	query := `INSERT INTO users(email, username, password)VALUES ($1, TRIM(LOWER($2)), $3)`
	_, err := database.Todo.Exec(query, email, username, password)
	if err != nil {
		return nil
	}
	return err
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	query := `SELECT id, password FROM users WHERE email = $1 AND archived_at IS NULL`
	err := database.Todo.Get(&user, query, email)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func CreateUserSession(user_id string, expiresAt time.Time) (string, error) {
	var sessionID string
	query := `INSERT INTO user_session(user_id, expires_at)VALUES ($1, $2)`
	err := database.Todo.Get(&sessionID, query, user_id, expiresAt)
	if err != nil {
		return "", err
	}
	return sessionID, nil
}

func DeleteUserSession(sessionID string) error {
	query := `UPDATE user_session SET archived_at = NOW() WHERE user_id = $1`
	_, err := database.Todo.Exec(query, sessionID)
	return err
}
