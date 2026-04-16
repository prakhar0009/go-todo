package dbHelper

import (
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

//func CreateUserSession(email, username, password string) (string, error) {
//
//}
//
//func DeleteUserSession(email string) error {
//
//}
