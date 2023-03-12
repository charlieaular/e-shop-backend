package repositories

import (
	"gorm.io/gorm"

	"e-shop-backend/src/models"
)

type CartRepo interface {
	CreateCart() (models.Cart, error)
	CurrentActiveCart() (models.Cart, error)
	UpdateCartPrice() error
	DesactivateCart() error
	PayCart() error
}

type cartRepo struct {
	DB *gorm.DB
}

func NewCartRepo(db *gorm.DB) CartRepo {
	return &cartRepo{DB: db}
}

func (cartRepo *cartRepo) CreateCart() (models.Cart, error) {
	newCart := models.Cart{
		Status: "ACTIVE",
	}

	result := cartRepo.DB.Create(&newCart)

	if result.Error != nil {
		return models.Cart{}, result.Error
	}

	return newCart, nil
}

func (cartRepo *cartRepo) CurrentActiveCart() (models.Cart, error) {
	var currentCart models.Cart

	result := cartRepo.DB.Model(models.Cart{}).Where("status = 'ACTIVE'").First(&currentCart)

	if result.Error != nil {
		return models.Cart{}, result.Error
	}

	return currentCart, nil

}
func (cartRepo *cartRepo) UpdateCartPrice() error {

	result := cartRepo.DB.Exec("UPDATE carts set total_price = (SELECT SUM(price) FROM cart_products WHERE cart_products.cart_id = carts.id) WHERE status = 'ACTIVE'")

	if result.Error != nil {
		return result.Error
	}
	return nil

}

func (cartRepo *cartRepo) DesactivateCart() error {
	return cartRepo.DB.Exec("UPDATE carts set status = 'DESACTIVED'").Error
}

func (cartRepo *cartRepo) PayCart() error {
	return cartRepo.DB.Model(models.Cart{}).Where("status = 'Active'").Update("status", "PAYED").Error
}
