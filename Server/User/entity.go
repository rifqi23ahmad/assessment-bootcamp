package user
import (
	"time"
)

type User struct {
	ID int gorm "primaryKey"
	FullName string
	Address string
	Email string
	Password string
}
