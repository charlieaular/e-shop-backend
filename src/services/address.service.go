package services

import (
	"e-shop-backend/src/models"
	"e-shop-backend/src/repositories"
)

type AddressService interface {
	CreateAddress(models.Address) (models.Address, error)
	GetAllAddresses(userId int) ([]models.Address, error)
	UpdateAddress(newAddress models.Address) (models.Address, error)
	DeleteAddress(id string) error
}

type addressService struct {
	repo repositories.AddressRepo
}

func NewAddressService(repo repositories.AddressRepo) AddressService {
	return &addressService{repo: repo}
}

func (s addressService) CreateAddress(newAddress models.Address) (models.Address, error) {
	return s.repo.CreateAddress(newAddress)
}

func (s addressService) GetAllAddresses(userId int) ([]models.Address, error) {
	return s.repo.GetAllAddresses(userId)
}

func (s addressService) UpdateAddress(newAddress models.Address) (models.Address, error) {
	return s.repo.UpdateAddress(newAddress)
}

func (s addressService) DeleteAddress(id string) error {
	return s.repo.DeleteAddress(id)
}
