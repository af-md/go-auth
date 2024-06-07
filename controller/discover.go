package controller

import (
	"cmp"
	"encoding/json"
	"go-auth/model"
	"log"
	"math/rand"
	"net/http"
	"slices"
	"strconv"

	"gorm.io/gorm"
)

type Result struct {
	User           model.User `json:"user"`
	DistanceFromMe int        `json:"distanceFromMe"`
}

type DiscoverResponse struct {
	Results []Result `json:"results"`
}

func Discover(w http.ResponseWriter, r *http.Request) {
	// return hello world back to the request
	// let get all the profiles from the user tables
	db := r.Context().Value("db").(*gorm.DB)
	email := r.Context().Value("email").(string)

	userLocation, err := getUserLocationByEmail(db, email)
	if err != nil {
		log.Print("discover: error getting user location: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Print("discover: user location: ", userLocation)

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

	usersSorted := sortByLocation(users, userLocation)

	result := make([]Result, 0)

	for _, v := range usersSorted {
		result = append(result, Result{
			User:           v,
			DistanceFromMe: v.Location,
		})
	}

	// filter based on swipe history
	d := DiscoverResponse{
		Results: result,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(d)
}

func getUserLocationByEmail(db *gorm.DB, email string) (string, error) {
	user := model.User{}
	result := db.Table("users").Where("email = ?", email).First(&user)
	if result.Error != nil {
		return "", result.Error
	}

	return strconv.Itoa(user.Location), nil
}

func sortByLocation(users []model.User, userLocation string) []model.User {

	// userLocation acts as minimum for random value
	userLocationConv, _ := strconv.Atoi(userLocation)
	max := 1000
	usersWithLocation := make([]model.User, 0)
	for _, v := range users {
		// create a random number that doesnt include the user location
		// if the random number is greater than the user location then we move the user to the right
		randomLocation := rand.Intn(max - userLocationConv + 1)
		// add minimum to the random location to make sure is more than 100
		randomLocation += userLocationConv
		v.Location = randomLocation
		usersWithLocation = append(usersWithLocation, v)
	}

	slices.SortFunc(usersWithLocation, func(a, b model.User) int {
		return cmp.Compare(a.Location, b.Location)
	})

	return usersWithLocation
}
