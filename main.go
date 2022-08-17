package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/priyanka19697/popcorn-be/database"
	"github.com/priyanka19697/popcorn-be/models"
	"github.com/priyanka19697/popcorn-be/routes"
)

func main() {
	db := database.Init()
	db.AutoMigrate(&models.Movie{}, &models.User{}, &models.Favorite{})
	// db.Model(&models.Favorite{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade")
	// db.Model(&models.Favorite{}).AddForeignKey("movie_id", "movies(id)", "cascade", "cascade")

	r := mux.NewRouter()

	// headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	// originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	// methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	// err := http.ListenAndServe("localhost:9010", handlers.CORS(originsOk, headersOk, methodsOk)(r))

	routes.RegisterMovieRoutes(r)
	http.Handle("/", r)
	// err := http.ListenAndServe("localhost:9010", r)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "HEAD", "POST", "PUT", "OPTIONS"},
	})

	handler := c.Handler(r)
	log.Fatal(http.ListenAndServe(":9010", handler))

	// if err != nil {
	// 	fmt.Println(err, "problem serving")
	// }

}
