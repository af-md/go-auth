// write a simple hello world server

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"go-auth/controller"
	"go-auth/services"

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

	router.Use(withDb(db))

	router.HandleFunc("/user/create", controller.CreateRandomUser)

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




// next steps:
// return the right result object
// package this with mysql and go application on docker container
// will other people get the right tables? will they get the data you have inside? what else will they get?
