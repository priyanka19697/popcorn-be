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

func CreateFavorite(favorite Favorite) (Favorite, error) {
	db := database.GetDB()
	var existingFavorite Favorite
	db.Where("user_id=? AND movie_id=?", favorite.UserID, favorite.MovieID).Find(&existingFavorite)
	result := db.Create(&favorite)
	return favorite, result.Error

}

func DeleteFavorite(userId int64, movieId int64) Favorite {
	db := database.GetDB()
	var favorite Favorite
	favorite, _ = FindFavorite(userId, movieId)
	var deletedFavorite Favorite
	db.Where("user_id=? AND movie_id=?", favorite.UserID, favorite.MovieID).Delete(&deletedFavorite)
	return deletedFavorite
}

func FindFavorite(userId int64, movieId int64) (Favorite, error) {
	db := database.GetDB()
	var favorite Favorite
	resultrows := db.Where("user_id=? AND movie_id=?", userId, movieId).Find(&favorite).RowsAffected
	// db.Preload("Movie").Find(&favorites)
	if resultrows == 0 {
		err := gorm.ErrRecordNotFound
		return favorite, err
	}
	return favorite, nil
}
