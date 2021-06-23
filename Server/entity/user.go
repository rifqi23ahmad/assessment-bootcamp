package entity

import "time"

type User struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	FullName    string    `json:"full_name"`
	Address     string    `json:"address"`
	Email       string    `json:"email"`
	Password    string    `json:"-"`
	PasswordMan []PassMan `json:"password_list"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserInput struct {
	FullName string `json:"full_name"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type PassManUserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserInput struct {
	FullName string `json:"full_name"`
	Address  string `json:"address"`
}
