package entity

import "time"

type PassMan struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	UserID    int       `json:"user_id"`
	Website   string    `json:"website"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type InputPassword struct {
	Website  string `json:"website"`
	Password string `json:"password"`
}
