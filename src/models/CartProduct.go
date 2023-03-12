package models

import (
	"gorm.io/gorm"
)

type CartProduct struct {
	gorm.Model
	CartID      int
	SizeID      int
	ColorID     int
	Name        string
	Image       string
	Description string
	Price       int
	Quantity    int
}
