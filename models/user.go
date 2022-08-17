package models

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/priyanka19697/popcorn-be/database"
)

type User struct {
	gorm.Model
	Name      string     `json:"username"`
	Email     string     `gorm:"type:varchar(100);unique_index" json:"email"`
	Password  string     `json:"password"`
	Favorites []Favorite `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func CreateUser(user User) error {
	db := database.GetDB()
	result := db.Create(&user)
	return result.Error
}

func GetAllUsers() []User {
	db := database.GetDB()
	var Users []User
	// db.Find(&Users)
	db.Model(&User{}).Preload("Favorites").Find(&Users)
	return Users
}

func GetUser(Id int64) (User, error) {
	db := database.GetDB()
	var getUser User
	result := db.Where("ID=?", Id).Find(&getUser)
	return getUser, result.Error
}

func ToggleFavorite(userId int64, movieId int64) []Favorite {
	var user User
	var movie Movie

	movie, _ = GetMovieById(movieId)
	user, _ = GetUser(userId)

	favorite := Favorite{}
	favorite.UserID = user.ID
	favorite.MovieID = movie.ID
	favorite.Movie = movie

	_, err := FindFavorite(userId, movieId)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		CreateFavorite(favorite)

	} else if err == nil {
		DeleteFavorite(userId, movieId)
	}
	result := ShowFavorites(userId)
	return result
}

func ShowFavorites(userId int64) []Favorite {
	db := database.GetDB()
	var favorites []Favorite
	// db.Model(&user).Association("Favorites").Find(&favorites)
	db.Where("user_id= ?", userId).Preload("Movie").Find(&favorites)
	return favorites
}
