package models

import (
	"github.com/jinzhu/gorm"
	"github.com/priyanka19697/popcorn-be/database"
)

type User struct {
	gorm.Model
	Name      string     `json:"username"`
	Email     string     `gorm:"type:varchar(100);unique_index" json:"email"`
	Password  string     `json:"password"`
	favorites []Favorite `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
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

func GetUser(Id int64) (*User, error) {
	db := database.GetDB()
	var getUser User
	result := db.Where("ID=?", Id).Find(&getUser)
	return &getUser, result.Error
}

func ToggleFavorite(userId int64, movieId int64) (Favorite, error) {
	db := database.GetDB()
	var getUser User
	var getMovie Movie
	db.Where("ID=?", userId).Find(&getUser)
	db.Where("ID=?", movieId).Find(&getMovie)
	favorite := Favorite{}
	favorite.UserID = getUser.ID
	favorite.MovieID = getMovie.ID
	favorite.Movie = getMovie
	result := db.Create(&favorite)
	return favorite, result.Error

}
