package repositories

import (
	"gorm.io/gorm"

	"e-shop-backend/src/models"
)

type CategoryRepo interface {
	GetAll() ([]models.Category, error)
	CreateCategory(models.Category) (models.Category, error)
	UpdateCategory(models.Category) (models.Category, error)
	DeleteCategory(string) error
}

type categoryRepo struct {
	DB *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) CategoryRepo {
	return &categoryRepo{DB: db}
}

func (r *categoryRepo) GetAll() ([]models.Category, error) {
	var categories []models.Category

	result := r.DB.Find(&categories)

	if result.Error != nil {
		return nil, result.Error
	}

	return categories, nil

}

func (r *categoryRepo) CreateCategory(newCategory models.Category) (models.Category, error) {
	err := r.DB.Create(&newCategory).Error

	if err != nil {
		return models.Category{}, err
	}

	return newCategory, nil

}

func (r *categoryRepo) UpdateCategory(newCategory models.Category) (models.Category, error) {
	err := r.DB.Updates(&newCategory).Error

	if err != nil {
		return models.Category{}, err
	}

	return newCategory, nil

}

func (r *categoryRepo) DeleteCategory(id string) error {
	err := r.DB.Delete(&models.Category{}, id).Error

	return err
}
