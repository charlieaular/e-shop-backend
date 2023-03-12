package repositories

import (
	"gorm.io/gorm"

	"e-shop-backend/src/models"
)

type AddressRepo interface {
	GetAllAddresses(userId int) ([]models.Address, error)
	CreateAddress(models.Address) (models.Address, error)
	UpdateAddress(newAddress models.Address) (models.Address, error)
	DeleteAddress(id string) error
}

type addressRepo struct {
	DB *gorm.DB
}

func NewAddressRepo(DB *gorm.DB) AddressRepo {
	return &addressRepo{
		DB: DB,
	}
}

func (r addressRepo) CreateAddress(newAddress models.Address) (models.Address, error) {
	err := r.DB.Create(&newAddress).Error

	if err != nil {
		return models.Address{}, err
	}

	return newAddress, nil
}

func (r addressRepo) GetAllAddresses(userId int) ([]models.Address, error) {

	var allAddresses []models.Address

	err := r.DB.Model(models.Address{}).Where("user_id = ?", userId).Find(&allAddresses).Error

	if err != nil {
		return nil, err
	}

	return allAddresses, nil

}

func (r addressRepo) UpdateAddress(newAddress models.Address) (models.Address, error) {
	err := r.DB.Updates(&newAddress).Error

	if err != nil {
		return models.Address{}, err
	}

	return newAddress, nil

}

func (r addressRepo) DeleteAddress(id string) error {
	err := r.DB.Delete(&models.Address{}, id).Error

	return err
}
