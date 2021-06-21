package user

import (
	"os/user"

	"gorm.io/gorm"
)

type Repository interface{
	UserGetAll() ([] User, error)
	UserCreated(user User) (User, error)
	UserbyID(userID string) (User, error)
}

type repository struct {
	db *gorm.DB

}
func NewRepository (db *gorm.DB) *repository {
	return &repository{db: }
}
func (r *rerepository) UserGetAll() ([] User, error) {
	var users [User]

	if err := r.db.Find(&users).Error; err != nil {
		return users, err
	}
}