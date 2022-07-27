package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/priyanka19697/popcorn-be/database"
	"github.com/priyanka19697/popcorn-be/models"
	"github.com/priyanka19697/popcorn-be/utils"

	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := utils.ParseBody(r, user)
	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
		res, _ := json.Marshal(resp)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return
	}
	resp := FindOne(user.Email, user.Password)
	res, _ := json.Marshal(resp)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func FindOne(email, password string) map[string]interface{} {
	user := &models.User{}
	db := database.GetDB()

	if err := db.Where("Email = ?", email).First(user).Error; err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Email address not found"}
		return resp
	}

	expiresAt := time.Now().Add(time.Minute * 100000).Unix()

	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword {
		var resp = map[string]interface{}{"status": false, "message": "Invalid login credentials. Please try again"}
		return resp
	}

	tokenObject := &models.Token{
		UserID: user.ID,
		Name:   user.Name,
		Email:  user.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenObject)

	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}

	var resp = map[string]interface{}{"status": false, "message": "logged in"}
	resp["token"] = tokenString //Store the token in the response
	resp["user"] = user
	return resp
}
