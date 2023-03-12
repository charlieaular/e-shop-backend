package models

import (
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	UserID  int
	Address string
	User    User
}
