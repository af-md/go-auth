package controller

import (
	"encoding/json"
	"go-auth/model"
	"io"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CustomResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

const tokenMessage = "User is logged in, make sure to save and attach the token to the next request as a Bearer token to show that the user is logged in. Without the token the request will be rejected."

func LoginUser(w http.ResponseWriter, r *http.Request) {
	// get db from context

	var loginRequest LoginRequest

	// decode request body into loginRequest
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Print("error reading body: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	err = json.Unmarshal(b, &loginRequest)

	log.Print("logged in: ", loginRequest)

	if err != nil {
		http.Error(w, "Invalid Payload Request", http.StatusBadRequest)
		return
	}

	// get user by email
	db := r.Context().Value("db").(*gorm.DB)
	var user model.User
	db.Where("email = ?", loginRequest.Email).First(&user)

	if user.ID == 0 {
		http.Error(w, "User not found, wrong email or password", http.StatusNotFound)
		return
	}

	log.Print("user found: ", user)

	// check password
	if user.Password != loginRequest.Password {
		http.Error(w, "User not found, wrong email or password", http.StatusNotFound)
		return
	}

	// generate token
	token := GenerateToken(user.Email)
	log.Print("token generated: ", token)

	cr := CustomResponse{
		Message: tokenMessage,
		Token:   token,
	}

	// attach the token to the request to show that the user is logged in
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cr)
}

func GenerateToken(email string) string {
	// generate token
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = jwt.MapClaims{
		"email": email,
	}

	tokenString, err := token.SignedString([]byte("my_secret"))
	if err != nil {
		log.Print("error generating token: ", err)
	}

	return tokenString
}
