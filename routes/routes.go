package routes

import (
	"github.com/gorilla/mux"
	"github.com/priyanka19697/popcorn-be/auth"
	"github.com/priyanka19697/popcorn-be/controllers"
)

var RegisterMovieRoutes = func(router *mux.Router) {
	router.HandleFunc("/movie", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/movie", controllers.ListMovies).Methods("GET")
	// router.HandleFunc("/movie/{movie}", controllers.GetMovieByTitle).Methods("GET")
	router.HandleFunc("/movie/{movieId}", controllers.GetMovieById).Methods("GET")

	router.HandleFunc("/movie/{movieId}", controllers.UpdateMovie).Methods("PUT")
	router.HandleFunc("/movie/{movieId}", controllers.DeleteMovie).Methods("DELETE")

	router.HandleFunc("/signup", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/login", controllers.LoginUser).Methods("POST")
	router.HandleFunc("/user", controllers.ListUsers).Methods("GET")

	// Auth route
	s := router.PathPrefix("/auth").Subrouter()
	s.Use(auth.JwtVerify)
	s.HandleFunc("/user", controllers.ListUsers).Methods("GET")
	s.HandleFunc("/user/{userId}", controllers.GetUser).Methods("GET")
}
