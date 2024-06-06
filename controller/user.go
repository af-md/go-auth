package controller

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"go-auth/model"
)

type Response struct {
	Result model.User `json:"result"`
}

func CreateRandomUser(w http.ResponseWriter, req *http.Request) {
	db, ok := req.Context().Value("db").(*gorm.DB)
	if !ok {
		log.Println("database not available in request context")
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	user, err := generateRandomUser(db)
	if err != nil {
		log.Printf("failed to create user: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	response := Response{Result: *user}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Printf("failed to marshal user: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}

func generateRandomUser(db *gorm.DB) (*model.User, error) {
	firstNames := []string{
		"Muhammad", "Ahmad", "Mahmoud", "Ibrahim", "Ali",
		"Yusuf", "Omar", "Hassan", "Amir", "Mustafa",
		"Fatima", "Aisha", "Zainab", "Layla", "Amina",
		"Noor", "Maryam", "Samira", "Rania", "Sadia",
	}

	lastNames := []string{
		"Khan", "Rahman", "Al-Mansour", "Farooq", "Malik",
		"Hussain", "Aziz", "Syed", "Patel", "Sheikh",
		"Anwar", "Yilmaz", "Hassan", "Iqbal", "Rashid",
		"El-Din", "Al-Hakim", "Saleh", "Bakri", "Siddiqi",
	}

	domains := []string{"gmail.com", "yahoo.com", "hotmail.com", "outlook.com", "example.com"}
	genders := []string{"Female", "Male"}

	firstName := firstNames[rand.Intn(len(firstNames))]
	lastName := lastNames[rand.Intn(len(lastNames))]
	name := firstName + " " + lastName

	email := generateRandomEmail(firstName, lastName, domains)
	password := generateRandomPassword()
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, err
	}

	age := uint(rand.Intn(30) + 18)
	gender := genders[rand.Intn(len(genders))]

	user := &model.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		Age:      age,
		Gender:   gender,
	}

	if err := db.Create(user).Error; err != nil {
		return nil, err
	}

	log.Printf("Created user with password: %s", string(password))
	return user, nil
}

func generateRandomEmail(firstName, lastName string, domains []string) string {
	username := strings.ToLower(firstName + "." + lastName)
	domain := domains[rand.Intn(len(domains))]
	return username + "@" + domain
}

func generateRandomPassword() []byte {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*"
	password := make([]byte, 8)
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}
	return password
}

func hashPassword(password []byte) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
