package models

type User struct {
	ID       string `db:"id" json:"id"`
	Email    string `db:"email" json:"email"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"-"`
	//CreatedAt time.Time `db:"created_at" json:"created_at"`
	//UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
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
