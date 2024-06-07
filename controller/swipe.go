package controller

import (
	"encoding/json"
	"go-auth/model"
	"io"
	"log"
	"net/http"

	"strconv"

	"gorm.io/gorm"
)

type SwipeRequest struct {
	Id         string `json:"id"`
	Preference string `json:"preference"`
}

func Swipe(w http.ResponseWriter, r *http.Request) {

	var swipeRequest SwipeRequest

	// decode request body into loginRequest
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Print("error reading body: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	err = json.Unmarshal(b, &swipeRequest)
	if err != nil {
		http.Error(w, "Invalid Payload Request", http.StatusBadRequest)
		return
	}

	if swipeRequest.Preference == "" && swipeRequest.Id == "" {
		http.Error(w, "Invalid Payload Request", http.StatusBadRequest)
		return
	}

	email := r.Context().Value("email").(string)
	db := r.Context().Value("db").(*gorm.DB)
	log.Print("email: ", email)
	// store the user and the swipe in the store

	// get the id of the user using the email from the db
	user := model.User{}
	db.Where("email = ?", email).First(&user)
	if user.ID == 0 {
		log.Print("The swiping user was not found with email: ", email)
		http.Error(w, "The swiping user was not found", http.StatusNotFound)
		return
	}

	conId, err := strconv.ParseUint(swipeRequest.Id, 10, 0)
	if err != nil {
		log.Print("error parsing id int uint: ", err)
		http.Error(w, "Invalid Payload Request", http.StatusBadRequest)
		return
	}

	swipe := model.Swipe{
		UserID:       user.ID,
		SwipedUserId: uint(conId),
	}
	// store the swipes in the db
	db.Table("swipe").Create(&swipe)

	// implement and check if there was a match

	log.Print("swiped: ", swipeRequest)

	// return match template

	// return hello world
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "swiped"}`))

}
