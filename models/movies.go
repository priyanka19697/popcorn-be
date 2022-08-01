package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/priyanka19697/popcorn-be/database"
)

type Movie struct {
	gorm.Model
	Title string `json:"title" validate:"required" gorm:"unique"`
	// Released   time.Time `json:"released" validate:"required" gorm:"type:datetime"`
	Released   string  `json:"released" validate:"required"`
	Genre      string  `json:"genre"`
	Type       string  `json:"type" validate:"required"`
	Plot       string  `json:"plot" validate:"required"`
	Poster     string  `json:"poster"`
	ImdbRating float32 `json:"rating" gorm:"default: 0.0"`
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

func GetMovieByTitle(name string) ([]Movie, *gorm.DB) {
	db := database.GetDB()
	var getMovie []Movie
	// db.Where("Title=?", name).Find(&getMovie)
	db.Where("Title LIKE ?", "%"+name+"%")
	return getMovie, db
}

func FindMoviesByYear(title string, year string) ([]Movie, error) {
	db := database.GetDB()
	var movies []Movie
	// datestring := "2022-08-01T15:37:46.811433349+05:30"
	// year = year + "-01" + "-01"
	// startYear, _ := time.Parse(year, datestring)
	// fmt.Println(year, startYear)
	// result := db.Where("Title LIKE ? AND year ", "%"+title+"%").Where("year >= ? AND year <= ?", ).Find(&movies)
	result := db.Where("Title LIKE ? AND released LIKE ?", "%"+title+"%", "%"+year+"%").Find(&movies)
	return movies, result.Error
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
		getMovie.Plot = movie.Plot
		getMovie.ImdbRating = movie.ImdbRating
		getMovie.Poster = movie.Poster
		getMovie.Genre = movie.Genre
		getMovie.Type = movie.Type
		getMovie.Released = movie.Released
		getMovie.UpdatedAt = time.Now()
		db.Save(&getMovie)
		fmt.Print(getMovie, "record after update")
		return getMovie, nil
	}
	return getMovie, result.Error
}
