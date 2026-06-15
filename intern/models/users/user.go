package users

import "time"

type User struct {
	ID           int       `json:"id"`
	FullName     string    `json:"fullname"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password"`
	CreatedAt    time.Time `json:"created_at"`
	Balance      int       `json:"balance"`
}
