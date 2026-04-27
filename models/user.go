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
func (r *UserRole) Scan(value interface{}) error {
	if value == nil {
		*r = RoleUser
		return nil
	}
	bv, ok := value.([]byte)
	if !ok {
		return errors.New("failed to scan UserRole")
	}
	*r = UserRole(string(bv))
	return nil
}

type User struct {
	ID       string   `db:"id" json:"id"`
	Email    string   `db:"email" json:"email"`
	Username string   `db:"username" json:"username"`
	Password string   `db:"password" json:"-"`
	Role     UserRole `db:"role" json:"role"`
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
