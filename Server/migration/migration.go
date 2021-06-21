package migration

import (
	"time"
)

type PassMan struct {
	ID        int `gorm: "primaryKey"`
	Website   string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	UserID   int `gorm: "primaryKey"`
	FullName string
	Address  string
	Email    string
	Password string
}
