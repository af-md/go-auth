package controller

import (
	"encoding/json"
	"go-auth/model"
	"log"
	"net/http"

	"gorm.io/gorm"
)

type DiscoverResponse struct {
	Results []model.User `json:"results"`
}

func Discover(w http.ResponseWriter, r *http.Request) {
	// return hello world back to the request
	// let get all the profiles from the user tables
	db := r.Context().Value("db").(*gorm.DB)

	q := r.URL.Query()

	// assume age and gender will always be in the query
	// assume that when age is queried we want to find all the users of that age and above
	age := q.Get("age")
	gender := q.Get("gender")

	var users []model.User
	db.Where("age >= ? AND gender = ?", age, gender).Find(&users)

	if len(users) == 0 {
		log.Print("discover: no users found when querying the database")
		log.Print("returning server error")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	log.Print("discover: amount of users found: ", len(users))

	// filter based on swipe history
	d := DiscoverResponse{
		Results: users,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(d)
}
