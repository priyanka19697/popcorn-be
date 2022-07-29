package models

import (
	"github.com/jinzhu/gorm"
	"github.com/priyanka19697/popcorn-be/database"
)

type Favorite struct {
	gorm.Model
	UserID  uint
	Movie   Movie `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	MovieID uint
}

func CreateFavorite(favorite Favorite) error {
	db := database.GetDB()
	result := db.Create(&favorite)
	return result.Error
}

func DeleteFavorite(userId int64, movieId int64) Favorite {
	db := database.GetDB()
	var favorite Favorite
	favorite, _ = FindFavorite(userId, movieId)
	var deletedFavorite Favorite
	db.Where("UserID=? AND MovieID=?", favorite.UserID, favorite.MovieID).Delete(deletedFavorite)
	return deletedFavorite
}

func FindFavorite(userId int64, movieId int64) (Favorite, error) {
	db := database.GetDB()
	var favorite Favorite
	result := db.Where("user_id=? AND movie_id=?", favorite.UserID, favorite.MovieID).Find(&favorite)
	return favorite, result.Error
}
