package validators

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/priyanka19697/popcorn-be/models"
	"github.com/priyanka19697/popcorn-be/utils"
)

var validate *validator.Validate

func CreateMovieValidator(r *http.Request) (models.Movie, error) {
	m := &models.Movie{}
	utils.ParseBody(r, m)

	validate = validator.New()
	err := validate.Struct(m)

	return *m, err
}
