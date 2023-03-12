package repositories

import (
	"gorm.io/gorm"

	"e-shop-backend/src/models"
)

type CartProductRepo interface {
	CreateCartProduct([]models.CartProduct) ([]models.CartProduct, error)
	GetCartProducts() ([]models.CartProduct, error)
	DeleteCartProduct(int) (bool, error)
}

type cartProductRepo struct {
	DB *gorm.DB
}

func NewCartProductRepo(db *gorm.DB) CartProductRepo {
	return &cartProductRepo{DB: db}
}

func (cartProductRepo *cartProductRepo) CreateCartProduct(newCartProducts []models.CartProduct) ([]models.CartProduct, error) {

	result := cartProductRepo.DB.Create(&newCartProducts)

	if result.Error != nil {
		return nil, result.Error
	}

	return newCartProducts, nil

}

func (cartProductRepo *cartProductRepo) GetCartProducts() ([]models.CartProduct, error) {

	var cartProducts []models.CartProduct

	err := cartProductRepo.DB.Model(models.CartProduct{}).Where("cart_id = (SELECT id FROM carts WHERE status = 'ACTIVE' LIMIT 1)").Find(&cartProducts).Error

	if err != nil {
		return nil, err
	}

	return cartProducts, nil
}

func (cartProductRepo *cartProductRepo) DeleteCartProduct(id int) (bool, error) {

	err := cartProductRepo.DB.Delete(&models.CartProduct{}, id).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
