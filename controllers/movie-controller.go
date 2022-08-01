package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/priyanka19697/popcorn-be/models"
	"github.com/priyanka19697/popcorn-be/utils"
	"github.com/priyanka19697/popcorn-be/validators"
)

var NewMovie models.Movie

func ListMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := models.GetAllMovies()

	var movieResponse utils.MovieResponse

	movieResponse.Data = movies
	if err != nil {
		movieResponse.Err = err.Error()
	}
	movieResponse.Err = "nil"

	res, _ := json.Marshal(movieResponse)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetMovieByTitle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movie := vars["movie"]
	movieRecord, _ := models.GetMovieByTitle(movie)
	fmt.Println(movieRecord)
	res, _ := json.Marshal(movieRecord)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetMovieById(w http.ResponseWriter, r *http.Request) {
	movieResponse := utils.MovieResponse{}

	vars := mux.Vars(r)
	movieId := vars["movieId"]

	ID, err := strconv.ParseInt(movieId, 0, 0)
	if err != nil {
		fmt.Println("parsing error")
	}
	movie, err := models.GetMovieById(ID)

	if err != nil {
		movieResponse.Err = err.Error()
	} else {
		movieResponse.Data = movie
		movieResponse.Err = "nil"
	}
	res, _ := json.Marshal(movieResponse)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func FindMoviesByYear(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("t")
	year := r.URL.Query().Get("y")
	fmt.Println(title, year)
	movies, _ := models.FindMoviesByYear(title, year)
	fmt.Println(movies)
	res, _ := json.Marshal(movies)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {

	movieResponse := utils.MovieResponse{}
	movie, err := validators.CreateMovieValidator(r)
	if err != nil {
		fmt.Println("Here error", err)
		movieResponse.Err = err.Error()
		res, _ := json.Marshal(movieResponse)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
		return
	}

	err2 := models.CreateMovie(movie)

	if err2 != nil {
		movieResponse.Err = err2.Error()
		res, _ := json.Marshal(movieResponse)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	} else {
		movieResponse.Err = ""
		movieResponse.Data = movie

		res, _ := json.Marshal(movieResponse)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}

}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieId := vars["movieId"]
	ID, err := strconv.ParseInt(movieId, 0, 0)
	if err != nil {
		fmt.Println("parsing error")
	}
	movie := models.DeleteMovie(ID)
	res, _ := json.Marshal(movie)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	movieResponse := utils.MovieResponse{}
	newMovieData, err := validators.CreateMovieValidator(r)
	vars := mux.Vars(r)
	movieId := vars["movieId"]
	ID, err := strconv.ParseInt(movieId, 0, 0)
	if err != nil {
		fmt.Println("parsing error")
	}
	movie, err := models.UpdateMovie(ID, newMovieData)

	if err != nil {
		movieResponse.Err = err.Error()
	} else {
		movieResponse.Data = movie
		movieResponse.Err = "nil"
	}
	res, _ := json.Marshal(movieResponse)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
