package models

import (
	"database/sql/driver"
	"errors"
)

// UserRole defines the custom type for our ENUM
type UserRole string

const (
	RoleUser  UserRole = "user"
	RoleAdmin UserRole = "admin"
)

// Value allows the database to store the string value
func (r UserRole) Value() (driver.Value, error) {
	return string(r), nil
}

// Scan allows the database to read the string into our custom type
// models/user.go

func (r *UserRole) Scan(value interface{}) error {
	if value == nil {
		*r = RoleUser
		return nil
	}

	switch v := value.(type) {
	case []byte:
		*r = UserRole(string(v))
	case string:
		*r = UserRole(v)
	default:
		return errors.New("failed to scan UserRole: incompatible type")
	}
	return nil
}

type User struct {
	ID        string   `db:"id" json:"id"`
	Email     string   `db:"email" json:"email"`
	Username  string   `db:"username" json:"username"`
	Password  string   `db:"password" json:"-"`
	Role      UserRole `db:"role" json:"role"`
	CreatedAt string   `db:"created_at" json:"created_at"`
}

type CreateUser struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUser struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
