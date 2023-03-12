package services

import (
	"e-shop-backend/src/models"
	"e-shop-backend/src/repositories"
)

type CategoryService interface {
	GetAll() ([]models.Category, error)
	CreateCategory(models.Category) (models.Category, error)
	UpdateCategory(models.Category) (models.Category, error)
	DeleteCategory(string) error
}

type categoryService struct {
	categoryRepo repositories.CategoryRepo
}

func NewCategoryService(categoryRepo repositories.CategoryRepo) CategoryService {
	return &categoryService{categoryRepo: categoryRepo}
}

func (categoryService *categoryService) GetAll() ([]models.Category, error) {
	return categoryService.categoryRepo.GetAll()
}

func (categoryService *categoryService) CreateCategory(newCategory models.Category) (models.Category, error) {
	return categoryService.categoryRepo.CreateCategory(newCategory)
}

func (categoryService *categoryService) UpdateCategory(newCategory models.Category) (models.Category, error) {
	return categoryService.categoryRepo.UpdateCategory(newCategory)
}

func (categoryService *categoryService) DeleteCategory(id string) error {
	return categoryService.categoryRepo.DeleteCategory(id)
}
