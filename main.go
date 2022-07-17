package main

import (
	"Dapperlabs_Challenge/endpointHandlers"
	"Dapperlabs_Challenge/models"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	//initialize DB connection pool
	err := models.InitDB()
	if err != nil {
		log.Fatalf("Error occured while attempting DB connection: %s", err)
	}

	router := mux.NewRouter()
	//Signup for new user and return valid JWT for authentication
	router.HandleFunc("/signup", endpointHandlers.Signup).Methods("POST")

	//Login as existing user and return JWT for authentication
	router.HandleFunc("/login", endpointHandlers.Login).Methods("POST")

	//Get list of all users
	router.HandleFunc("/users", endpointHandlers.GetUsers).Methods("GET")

	//Update firstname & lastname of current user
	router.HandleFunc("/users", endpointHandlers.UpdateUser).Methods("PUT")

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
