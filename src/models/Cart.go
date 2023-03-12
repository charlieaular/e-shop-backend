package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID     int
	Status     string
	TotalPrice int
	User       User
}
