package migration

import "time"

type User struct {
	ID        int `gorm:"primaryKey"`
	FullName  string
	Address   string
	Email     string
	Password  string
	PassMan   []PassMan `gorm:"foreignKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PassMan struct {
	ID        int `gorm:"primaryKey" `
	UserID    string
	Website   string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
