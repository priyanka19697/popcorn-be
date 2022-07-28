package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/priyanka19697/popcorn-be/database"
	"github.com/priyanka19697/popcorn-be/models"
	"github.com/priyanka19697/popcorn-be/routes"
)

func main() {
	db := database.Init()
	db.AutoMigrate(&models.Movie{}, &models.User{}, &models.Favorite{})

	r := mux.NewRouter()
	routes.RegisterMovieRoutes(r)
	http.Handle("/", r)
	err := http.ListenAndServe("localhost:9010", r)

	if err != nil {
		fmt.Println(err, "problem serving")
	}

}
