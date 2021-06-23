package user

import (
	"gorm.io/gorm"
	"assessment-bootcamp/entity"
)

type Repository interface {
	UserGetAll() ([]entity.User, error)
	UserCreated(userCreate entity.user User) (entity.User, error)
	UserbyID(UserbyID string) (User, error)
	UserFindEmail(email string) (User, error)
	UserDelete(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) UserGetiAll() ([]User, error) {
	var users []entity.User
	err := r.db.Find(&users).error
	if err != nil {
		return users
	}
	return users, nil

}
