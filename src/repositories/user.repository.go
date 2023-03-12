package repositories

import (
	"gorm.io/gorm"

	"e-shop-backend/src/models"
)

type UserRepo interface {
	GetUserByEmail(string) (models.User, error)
	GetUserById(int) (models.User, error)
	CreateClient(models.User) (models.User, error)
}

type userRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{DB: db}
}

func (userRepo *userRepo) GetUserByEmail(email string) (models.User, error) {
	var user models.User

	result := userRepo.DB.Model(models.User{}).Where("email = ?", email).First(&user)

	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil

}

func (userRepo *userRepo) GetUserById(id int) (models.User, error) {
	var user models.User

	result := userRepo.DB.Model(models.User{}).First(&user, id)

	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}

func (userRepo *userRepo) CreateClient(user models.User) (models.User, error) {

	result := userRepo.DB.Create(&user)

	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}
