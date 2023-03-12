package models

import (
	"gorm.io/gorm"
)

type Size struct {
	gorm.Model
	Name string
}
