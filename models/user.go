package models

import (
	"github.com/jinzhu/gorm"
	"github.com/priyanka19697/popcorn-be/database"
)

type User struct {
	gorm.Model
	Name     string `json:"username"`
	Email    string `gorm:"type:varchar(100);unique_index" json:"email"`
	Password string `json:"password"`
}

func CreateUser(user User) error {
	db := database.GetDB()
	result := db.Create(&user)
	return result.Error
}

func GetAllUsers() []User {
	db := database.GetDB()
	var Users []User
	db.Find(&Users)
	return Users
}
