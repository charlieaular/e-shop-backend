package models

import (
	"gorm.io/gorm"
)

type Color struct {
	gorm.Model
	Name  string
	Color string
}
