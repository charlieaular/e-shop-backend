package services

import (
	"e-shop-backend/src/models"
	"e-shop-backend/src/repositories"
)

type CartService interface {
	CreateCart() (models.Cart, error)
	CurrentActiveCart() (models.Cart, error)
	UpdateCartPrice() error
	DesactivateCart() error
	PayCart() error
}

type cartService struct {
	cartRepo repositories.CartRepo
}

func NewCartService(cartRepo repositories.CartRepo) CartService {
	return &cartService{cartRepo: cartRepo}
}

func (cartService *cartService) CreateCart() (models.Cart, error) {
	return cartService.cartRepo.CreateCart()
}

func (cartService *cartService) CurrentActiveCart() (models.Cart, error) {
	return cartService.cartRepo.CurrentActiveCart()
}

func (cartService *cartService) UpdateCartPrice() error {
	return cartService.cartRepo.UpdateCartPrice()
}

func (cartService *cartService) DesactivateCart() error {
	return cartService.cartRepo.DesactivateCart()
}

func (cartService *cartService) PayCart() error {
	return cartService.cartRepo.PayCart()
}
