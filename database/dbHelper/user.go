package dbHelper

import (
	"fmt"
	"time"

	"github.com/prakhar0009/go-todo/database"
	"github.com/prakhar0009/go-todo/models"
)

func IsUserExist(username string) (bool, error) {
	var exist bool
	query := `SELECT EXISTS (SELECT * FROM users WHERE username = TRIM(LOWER($1)) AND archived_at IS NULL)`
	err := database.Todo.Get(&exist, query, username)
	return exist, err
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
		"SELECT id, email, username, password, role FROM users WHERE email=$1 AND archived_at IS NULL",
		email,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func CreateUserSession(userID string, expiry time.Time) (string, error) {
	var sessionID string

	query := "INSERT INTO user_session (user_id, expires_at) VALUES ($1, $2) RETURNING id"

	err := database.Todo.Get(&sessionID, query, userID, expiry)

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
func GetUserBySession(sessionID string) (string, error) {
	query := `select user_id from user_session where id=$1 and archived_at IS NULL`
	var userId string
	err := database.Todo.Get(&userId, query, sessionID)
	if err != nil {
		return "", err
	}
	return userId, nil
}

func UpdateUserRole(userID string, newRole models.UserRole) error {
	query := `
				UPDATE users 
				SET role = $1
				WHERE id = $2
				AND archived_at IS NULL
				`
	_, err := database.Todo.Exec(query, newRole, userID)
	return err
}
