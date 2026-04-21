package models

import "time"

type Todo struct {
	ID           string `json:"id" db:"id"`
	UserID       string `json:"user_id" db:"user_id"`
	Title        string `json:"title" db:"title"`
	Description  string `json:"description" db:"description"`
	IsCompleted  bool   `json:"is_completed" db:"is_completed"`
	IsIncomplete bool   `json:"is_incomplete" db:"is_incomplete"`
	IsPending    bool   `json:"is_pending" db:"is_pending"`
	ExpiresAt    string `json:"expires_at" db:"expires_at"`
	CreatedAt    string `json:"created_at" db:"created_at"`
	//ArchivedAt  string `json:"archived_at" db:"archived_at,omitempty"`
}

type CreateTodo struct {
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	ExpiresAt   string `json:"expires_at"`
}

type UpdateTodo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted *bool  `json:"is_completed"` // used pointer here to distinguish between 'false' and 'not_sent'
}

func (t *Todo) SyncStatus() {
	if t.IsCompleted {
		t.IsIncomplete = false
		t.IsPending = false
		return
	}

	expiry, err := time.Parse(time.RFC3339, t.ExpiresAt)
	if err != nil {
		t.IsIncomplete = true
		t.IsPending = false
		return
	}

	if time.Now().After(expiry) {
		t.IsPending = true
		t.IsIncomplete = false
	} else {
		t.IsPending = false
		t.IsIncomplete = true
	}
}
