package models

import (
	"gorm.io/gorm"
)

type ProductCategory struct {
	gorm.Model
	ProductID  int
	CategoryID int
}
