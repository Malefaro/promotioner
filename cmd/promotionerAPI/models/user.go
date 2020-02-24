package models

type User struct {
	UserID   int64    `json:"user_id"`
	Email    string `json:"email"`
	Password string `json:"password, omitempty"`
}
