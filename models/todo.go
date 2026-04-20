package models

type Todo struct {
	ID          string `json:"id" db:"id"`
	UserID      string `json:"user_id" db:"user_id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	IsCompleted bool   `json:"is_completed" db:"is_completed"`
	ExpiresAt   string `json:"expires_at" db:"expires_at"`
	CreatedAt   string `json:"created_at" db:"created_at"`
	//ArchivedAt  string `json:"archived_at" db:"archived_at,omitempty"`
}
