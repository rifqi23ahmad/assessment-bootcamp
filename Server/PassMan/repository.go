package passman

import (
	"assessment-bootcamp/entity"

	"gorm.io/gorm"
)

type PassManRepository interface {
	PassGetPass() ([]entity.PassMan, error)
	PassCreated(Passman entity.PassMan) (entity.PassMan, error)
	PassDelete(id string) (string, error)
	UpdatePassword(id string, dataUpdate map[string]interface{}) (entity.PassMan, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) PassGetPass(id string) (entity.PassMan, error) {
	var Passman entity.PassMan

	if err := r.db.Where("id = ?", id).Find(&Passman).Error; err != nil {
		return Passman, err
	}

	return Passman, nil
}

func (r *Repository) PassCreated(Passman entity.PassMan) (entity.PassMan, error) {
	if err := r.db.Create(&Passman).Error; err != nil {
		return Passman, err
	}

	return Passman, nil

}

func (r *Repository) PassDelete(id string) (string, error) {
	if err := r.db.Where("id = ?", id).Delete(&entity.PassMan{}).Error; err != nil {
		return "error", err
	}

	return "success", nil
}

func (r *Repository) UpdatePassword(id string, dataUpdate map[string]interface{}) (entity.PassMan, error) {

	var PassMan entity.PassMan
	if err := r.db.Model(&PassMan).Where("id = ?", id).Updates(dataUpdate).Error; err != nil {
		return PassMan, err
	}

	if err := r.db.Where("id = ?", id).Find(&PassMan).Error; err != nil {
		return PassMan, err
	}

	return PassMan, nil

}
