package services

import (
	"e-shop-backend/src/models"
	"e-shop-backend/src/repositories"
)

type ProductService interface {
	GetProductsByCategory(int) ([]models.Product, error)
	GetById(int) (models.Product, error)
	GetByIds([]int) ([]models.Product, error)
}

type productService struct {
	productRepo repositories.ProductRepo
}

func NewProductService(productRepo repositories.ProductRepo) ProductService {
	return &productService{productRepo: productRepo}
}

func (productService *productService) GetProductsByCategory(category int) ([]models.Product, error) {
	return productService.productRepo.GetProductsByCategory(category)
}

func (productService *productService) GetById(id int) (models.Product, error) {
	return productService.productRepo.GetById(id)
}

func (productService *productService) GetByIds(ids []int) ([]models.Product, error) {
	return productService.productRepo.GetByIds(ids)
}
