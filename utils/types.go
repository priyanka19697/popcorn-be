package utils

type MovieResponse struct {
	Data interface{} `json:"data"`
	Err  string      `json:"err"`
}

type UserResponse struct {
	Data interface{} `json:"data"`
	Err  string      `json:"err"`
}

type Exception struct {
	Message string `json:"message"`
}

type MovieType string

const (
	Movie  MovieType = "movie"
	Series MovieType = "series"
)
