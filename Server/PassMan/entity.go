package passman

import (
	"time"
)

type PassMan struct {
	ID int gorm "primaryKey"
	Website string
	Email string 
	Password string
	CreatedAt time.time 
	UpdatedAt time.time
}