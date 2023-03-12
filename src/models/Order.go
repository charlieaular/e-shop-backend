package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID int
	CartID int
	Status string
	User   User
	Cart   Cart
}
