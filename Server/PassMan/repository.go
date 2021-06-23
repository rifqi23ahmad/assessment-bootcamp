package password

import (
	"assessment-bootcamp/Server/entity"

	"gorm.io/gorm"
)

type PassManRepository interface {
	RepCreatePass(password entity.PassMan) (entity.PassMan, error)
	RepGetPass(id string) (entity.PassMan, error)
	RepUpdatePass(id string, dataUpdate map[string]interface{}) (entity.PassMan, error)
	RepDelPass(id string) (string, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) RepCreatePass(password entity.PassMan) (entity.PassMan, error) {
	if err := r.db.Create(&password).Error; err != nil {
		return password, err
	}

	return password, nil

}

func (r *Repository) RepGetPass(id string) (entity.PassMan, error) {
	var password entity.PassMan

	if err := r.db.Where("id = ?", id).Find(&password).Error; err != nil {
		return password, err
	}

	return password, nil
}

func (r *Repository) RepUpdatePass(id string, dataUpdate map[string]interface{}) (entity.PassMan, error) {
	var password entity.PassMan

	if err := r.db.Model(&password).Where("id = ?", id).Updates(dataUpdate).Error; err != nil {
		return password, err
	}

	if err := r.db.Where("id = ?", id).Find(&password).Error; err != nil {
		return password, err
	}

	return password, nil
}

func (r *Repository) RepDelPass(id string) (string, error) {
	if err := r.db.Where("id = ?", id).Delete(&entity.PassMan{}).Error; err != nil {
		return "error", err
	}

	return "success", nil
}
