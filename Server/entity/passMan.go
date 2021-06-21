package passman

import "time"

type PassMan struct {
	ID        int `gorm: "primaryKey" json:"id"`
	Website   string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
