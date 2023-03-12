package services

import (
	"e-shop-backend/src/models"
	"e-shop-backend/src/repositories"
)

type CartProductService interface {
	CreateCartProduct([]models.CartProduct) ([]models.CartProduct, error)
	GetCartProducts() ([]models.CartProduct, error)
	DeleteCartProduct(int) (bool, error)
}

type cartProductService struct {
	cartProductRepo repositories.CartProductRepo
}

func NewCartProductService(cartProductRepo repositories.CartProductRepo) CartProductService {
	return &cartProductService{cartProductRepo: cartProductRepo}
}

func (cartProductService *cartProductService) CreateCartProduct(newCartProducts []models.CartProduct) ([]models.CartProduct, error) {
	return cartProductService.cartProductRepo.CreateCartProduct(newCartProducts)
}

func (cartProductService *cartProductService) GetCartProducts() ([]models.CartProduct, error) {
	return cartProductService.cartProductRepo.GetCartProducts()
}

func (cartProductService *cartProductService) DeleteCartProduct(id int) (bool, error) {
	return cartProductService.cartProductRepo.DeleteCartProduct(id)
}
