package controllers

type MovieResponse struct {
	Data interface{} `json:"data"`
	Err  string      `json:"err"`
}

type UserResponse struct {
	Data interface{} `json:"data"`
	Err  string      `json:"err"`
}
