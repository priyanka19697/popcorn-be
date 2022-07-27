package models

import (
	"github.com/jinzhu/gorm"
)

type Favorite struct {
	gorm.Model
	Name    string
	UserId  uint   `gorm:"foreignkey:UserId"`
	Movie   *Movie `gorm:"foreignkey:MovieId"`
	MovieId uint
}
