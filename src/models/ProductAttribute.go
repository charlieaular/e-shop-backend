package models

import (
	"gorm.io/gorm"
)

type ProductAttribute struct {
	gorm.Model
	ProductID int
	SizeID    string `gorm:"default:[]"`
	ColorID   string `gorm:"default:[]"`
}
