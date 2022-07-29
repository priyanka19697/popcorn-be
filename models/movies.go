package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/priyanka19697/popcorn-be/database"
)

type Movie struct {
	gorm.Model
	Title       string    `json:"title" validate:"required" gorm:"unique"`
	Year        time.Time `json:"year" validate:"required" gorm:"type:datetime"`
	Rating      float32   `json:"rating" gorm:"default: 0.0"`
	Url         string    `json:"url" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Type        string    `json:"type" validate:"required"`
	PosterURL   string    `json:"posterURL"`
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

func GetMovieById(Id int64) (Movie, error) {
	db := database.GetDB()
	var movie Movie
	result := db.Where("ID=?", Id).Find(&movie)
	return movie, result.Error
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

func UpdateMovie(Id int64, movie Movie) (Movie, error) {
	db := database.GetDB()
	var getMovie Movie
	result := db.Where("ID=?", Id).Find(&getMovie)

	fmt.Printf("%+v to be updated with", movie)
	fmt.Printf("%v %v", result, result.Error)

	if result.Error == nil {
		getMovie.Title = movie.Title
		getMovie.Description = movie.Description
		getMovie.Rating = movie.Rating
		getMovie.Url = movie.Url
		getMovie.PosterURL = movie.PosterURL
		getMovie.Type = movie.Type
		getMovie.Year = movie.Year
		getMovie.UpdatedAt = time.Now()
		db.Save(&getMovie)
		fmt.Print(getMovie, "record after update")
		return getMovie, nil
	}
	return getMovie, result.Error
}
