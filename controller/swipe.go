package controller

import (
	"encoding/json"
	"errors"
	"fmt"
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

type Matched struct {
	Matched bool `json:"matched"`
	MatchId uint `json:"matchID"`
}

type Match struct {
	Matched bool `json:"matched"`
}

type SwipeResponse struct {
	Result interface{} `json:"result"`
}

func Swipe(w http.ResponseWriter, r *http.Request) {

	var swipeRequest SwipeRequest

	// parse request body
	err := parseRequestBody(r, &swipeRequest)
	if err != nil {
		log.Print("error parsing request body: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// extract the email from the context
	email := r.Context().Value("email").(string)
	db := r.Context().Value("db").(*gorm.DB)
	log.Print("email: ", email)

	// get the id of the user using the email from the db
	user := model.User{}
	db.Where("email = ?", email).First(&user)
	if user.ID == 0 {
		log.Print("The swiping user was not found with email: ", email)
		http.Error(w, "The swiping user was not found", http.StatusNotFound)
		return
	}

	user, err = getUserByEmail(db, email)
	if err != nil {
		log.Print("error getting user by email: ", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
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
		Preference:   swipeRequest.Preference,
	}

	// if the record is not found, the library logs the error automatically for me
	// check if the swipe already exists
	checkSwipe, err := checkSwipeExists(db, swipe)
	if err != nil {
		log.Print("error checking if swipe exists: ", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if checkSwipe.ID == 0 {
		// add the swipe to the db
		err = addSwipeToDB(db, swipe)
		if err != nil {
			log.Print("error adding swipe to db: ", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}

	// check if the user has a match with the swiped user
	matchCheck := model.Swipe{}
	// querying for a match
	result := db.Table("swipe").Where("UserId = ? AND SwipedUserId = ? AND Preference = ?", swipe.SwipedUserId, swipe.UserID, "Yes").First(&matchCheck)
	if result.Error != nil {
		log.Print("error checking for match: ", result.Error)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(SwipeResponse{Result: Match{Matched: false}})
		return
	}
	log.Print("Match found")

	match := model.Match{
		UserId1: swipe.UserID,
		UserId2: swipe.SwipedUserId,
	}

	// check if the match already exists
	matchCheckInMatchTable := model.Match{}
	match, err = checkMatch(db, matchCheckInMatchTable)
	if err != nil {
		log.Print("error checking if match exists: ", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if match.ID != 0 {
		log.Print("match record already exists for user with id: ", match.UserId1, " and swiped user id: ", match.UserId2)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(SwipeResponse{Result: Matched{Matched: true, MatchId: match.ID}})
		return
	}

	// if the user was matched with the swiped user, store in match table
	result = db.Table("matches").Create(&match)
	if result.Error != nil {
		log.Print("error creating match: ", result.Error)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// fetch the the id of the match from the match table using the user id and swiped user id
	result = db.Table("matches").Where("UserId1 = ? AND UserId2 = ?", match.UserId1, match.UserId2).First(&match)
	if result.Error != nil {
		log.Print("error fetching match: ", result.Error)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// return the match id to the user with matched: true
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(SwipeResponse{Result: Matched{Matched: true, MatchId: match.ID}})
}

func parseRequestBody(r *http.Request, s *SwipeRequest) error {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, s)
	if err != nil {
		return err
	}

	if s.Preference == "" && s.Id == "" {
		return fmt.Errorf("empty strings, invalid payload request")
	}

	return nil
}

func getUserByEmail(db *gorm.DB, email string) (model.User, error) {
	user := model.User{}
	result := db.Table("users").Where("email = ?", email).First(&user)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func checkSwipeExists(db *gorm.DB, swipe model.Swipe) (model.Swipe, error) {
	checkSwipe := model.Swipe{}
	result := db.Table("swipe").Where("UserId = ? AND SwipedUserId = ?", swipe.UserID, swipe.SwipedUserId).First(&checkSwipe)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// Record not found, return an empty swipe and nil error
			return model.Swipe{}, nil
		}
		// Other error occurred, return the error
		return model.Swipe{}, result.Error
	}
	return checkSwipe, nil
}

func addSwipeToDB(db *gorm.DB, swipe model.Swipe) error {
	result := db.Table("swipe").Create(&swipe)
	if result.Error != nil {
		return result.Error
	}
	return nil
}


// are there good consistencies passing models ?? i mean &
func checkMatch(db *gorm.DB, match model.Match) (model.Match, error) {
	checkMatch := model.Match{}
	result := db.Table("matches").Where("UserId1 = ? AND UserId2 = ?", match.UserId1, match.UserId2).First(&checkMatch)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// Record not found, return an empty swipe and nil error
			return model.Match{}, nil
		}
		// Other error occurred, return the error
		return model.Match{}, result.Error
	}

	return checkMatch, nil
}
