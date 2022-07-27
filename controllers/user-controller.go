package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/priyanka19697/popcorn-be/models"
	"github.com/priyanka19697/popcorn-be/utils"

	"golang.org/x/crypto/bcrypt"
)

type ErrorResponse struct {
	Err string
}

var NewUser models.User

func CreateUser(w http.ResponseWriter, r *http.Request) {

	userResponse := UserResponse{}

	user := &models.User{}
	utils.ParseBody(r, user)

	fmt.Printf("%+v", user)

	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		err := ErrorResponse{
			Err: "Password Encryption failed",
		}
		json.NewEncoder(w).Encode(err)
	}

	user.Password = string(pass)
	err2 := models.CreateUser(*user)

	if err != nil {
		userResponse.Err = err2.Error()
		res, _ := json.Marshal(userResponse)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	} else {
		userResponse.Err = ""
		userResponse.Data = user
	}
	res, _ := json.Marshal(user)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	users := models.GetAllUsers()
	res, _ := json.Marshal(users)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
