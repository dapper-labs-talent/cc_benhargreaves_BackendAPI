package main

import (
	"Dapperlabs_Challenge/endpointHandlers"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	//Signup for new user and return valid JWT for authentication
	router.HandleFunc("/signup", endpointHandlers.Signup).Methods("POST")

	//Login as existing user and return JWT for authentication
	router.HandleFunc("/login", endpointHandlers.Login).Methods("POST")

	//Get list of all users
	router.HandleFunc("/users", endpointHandlers.GetUser).Methods("GET")

	//Update firstname & lastname of current user
	router.HandleFunc("/users", endpointHandlers.UpdateUser).Methods("PUT")

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
