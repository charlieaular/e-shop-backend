package repositories

import (
	"gorm.io/gorm"

	"e-shop-backend/src/models"
)

type TagRepo interface {
	GetAllTags() ([]models.Tag, error)
	CreateTag(models.Tag) (models.Tag, error)
	UpdateTag(models.Tag) (models.Tag, error)
	DeleteTag(string) error
}

type tagRepo struct {
	DB *gorm.DB
}

func NewTagRepository(db *gorm.DB) TagRepo {
	return &tagRepo{DB: db}
}

func (r tagRepo) GetAllTags() ([]models.Tag, error) {
	var allTags []models.Tag

	err := r.DB.Model(models.Tag{}).Find(&allTags).Error

	if err != nil {
		return nil, err
	}

	return allTags, nil
}

func (r *tagRepo) CreateTag(newTag models.Tag) (models.Tag, error) {
	err := r.DB.Create(&newTag).Error

	if err != nil {
		return models.Tag{}, err
	}

	return newTag, nil

}

func (r *tagRepo) UpdateTag(newTag models.Tag) (models.Tag, error) {
	err := r.DB.Updates(&newTag).Error

	if err != nil {
		return models.Tag{}, err
	}

	return newTag, nil

}

func (r *tagRepo) DeleteTag(id string) error {
	err := r.DB.Delete(&models.Tag{}, id).Error

	return err
}
