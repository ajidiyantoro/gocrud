package user

import "time"

type User struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Gender      string    `json:"gender"`
	Dateofbirth string    `json:"dataofbirth"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	IsActived   int       `json:"is_actived"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetUser struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
