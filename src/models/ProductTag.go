package models

import (
	"gorm.io/gorm"
)

type ProductTag struct {
	gorm.Model
	ProductID int
	TagsID    string `gorm:"default:[]"`
}
