package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/priyanka19697/popcorn-be/database"
)

type Movie struct {
	gorm.Model
	Title       string  `json:"title" validate:"required" gorm:"unique"`
	Year        string  `json:"year" validate:"required"`
	Rating      float32 `json:"rating" gorm:"default: 0.0"`
	Url         string  `json:"url" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Type        string  `json:"type" validate:"required"`
	PosterURL   string  `json:"posterURL"`
}

func CreateMovie(movie Movie) error {
	db := database.GetDB()
	result := db.Create(&movie)
	return result.Error
}

func GetAllMovies() ([]Movie, error) {
	db := database.GetDB()
	var Movies []Movie

	result := db.Find(&Movies)

	if result.Error != nil {
		return nil, fmt.Errorf("there was a problem fetching data in GetAllMovies")
	}

	return Movies, nil
}

func GetMovieById(Id int64) (*Movie, *gorm.DB) {
	db := database.GetDB()
	var getMovie Movie
	db.Where("ID=?", Id).Find(&getMovie)
	return &getMovie, db
}

func GetMovieByTitle(name string) (*Movie, *gorm.DB) {
	db := database.GetDB()
	var getMovie Movie
	db.Where("Title=?", name).Find(&getMovie)
	return &getMovie, db
}

func DeleteMovie(Id int64) Movie {
	db := database.GetDB()
	var movie Movie
	db.Where("ID= ?", Id).Delete(movie)
	return movie
}
