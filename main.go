// write a simple hello world server

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"go-auth/controller"
	"go-auth/services"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	// orm
)

// {
// 	"result": {
// 	"id": <integer>,
// 	"email": <string>,
// 	"password": <string>,
// 	"name": <string>,
// 	"gender": <string>,
// 	"age": <integer>
// 	}
// }

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func main() {

	// create a db
	db, err := services.CreateDB()
	log.Print("db created")

	if err != nil {
		fmt.Println(err)
		fmt.Print("shutting down gracefully")
		os.Exit(1)
	}

	// create a new serve mux i.e. a new router
	router := mux.NewRouter()

	// attach the middleware
	router.Use(withDb(db))
	router.Use(withToken())

	router.HandleFunc("/user/create", controller.CreateRandomUser).Methods("GET")
	router.HandleFunc("/login", controller.LoginUser).Methods("POST")
	router.HandleFunc("/discover", controller.Discover).Methods("GET")
	router.HandleFunc("/swipe", controller.Swipe).Methods("POST")

	http.ListenAndServe(":8080", router)
}

func withDb(db *gorm.DB) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("middleware used")
			ctx := context.WithValue(r.Context(), "db", db)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func withToken() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// check if the request is for a login or create user and skip the token check

			if r.URL.Path == "/login" || r.URL.Path == "/user/create" {
				next.ServeHTTP(w, r)
				return
			}

			log.Printf("middleware used token")
			authHeader := r.Header.Get("Authorization")

			// Split the Authorization header to extract the JWT
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, "Invalid Authorization header", http.StatusUnauthorized)
				return
			}

			tokenString := parts[1]
			log.Print("token: ", tokenString)

			// Parse and validate the JWT
			token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
				// Provide the secret key used for signing the JWT
				return []byte("my_secret"), nil
			})

			if err != nil {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			if !token.Valid {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			// Extract the claims from the JWT
			claims, ok := token.Claims.(*Claims)
			if !ok {
				http.Error(w, "Invalid token claims", http.StatusUnauthorized)
				return
			}

			// Set the email in the request context for further use
			ctx := r.Context()
			ctx = context.WithValue(ctx, "email", claims.Email)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
