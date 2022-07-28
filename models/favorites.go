package models

import (
	"github.com/jinzhu/gorm"
)

type Favorite struct {
	gorm.Model
	UserID  uint
	Movie   Movie `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	MovieID uint
}
