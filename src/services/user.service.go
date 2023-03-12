package services

import (
	"e-shop-backend/src/models"
	"e-shop-backend/src/repositories"
)

type UserService interface {
	GetUserByEmail(string) (models.User, error)
	GetUserById(int) (models.User, error)
	CreateClient(models.User) (models.User, error)
}

type userService struct {
	userRepo repositories.UserRepo
}

func NewUserService(userRepo repositories.UserRepo) UserService {
	return &userService{userRepo: userRepo}
}

func (userService *userService) GetUserByEmail(email string) (models.User, error) {
	return userService.userRepo.GetUserByEmail(email)
}

func (userService *userService) GetUserById(id int) (models.User, error) {
	return userService.userRepo.GetUserById(id)
}

func (userService *userService) CreateClient(user models.User) (models.User, error) {
	return userService.userRepo.CreateClient(user)
}
