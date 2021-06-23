package user

import (
	"assessment-bootcamp/Server/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user entity.User) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
	UpdateUser(id string, dataUpdate map[string]interface{}) (entity.User, error)
	GetOneUser(id string) (entity.User, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) CreateUser(user entity.User) (entity.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *Repository) FindByEmail(email string) (entity.User, error) {
	var user entity.User

	if err := r.db.Where("email = ?", email).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *Repository) GetOneUser(id string) (entity.User, error) {
	var user entity.User

	if err := r.db.Where("id = ?", id).Preload("PasswordList").Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *Repository) UpdateUser(id string, dataUpdate map[string]interface{}) (entity.User, error) {
	var user entity.User

	if err := r.db.Model(&user).Where("id = ?", id).Updates(dataUpdate).Error; err != nil {
		return user, err
	}

	if err := r.db.Where("id = ?", id).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
