package routes

import (
	"github.com/gorilla/mux"
	"github.com/priyanka19697/popcorn-be/controllers"
)

var RegisterMovieRoutes = func(router *mux.Router) {
	router.HandleFunc("/movie", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/movie", controllers.ListMovies).Methods("GET")
	router.HandleFunc("/movie/{movie}", controllers.GetMovieByTitle).Methods("GET")
	router.HandleFunc("/movie/{movieId}", controllers.UpdateMovie).Methods("PUT")
	router.HandleFunc("/movie/{movieID}", controllers.DeleteMovie).Methods("DELETE")

	router.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user", controllers.ListUsers).Methods("GET")

}
