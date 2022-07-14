package endpointHandlers

import (
	"Dapperlabs_Challenge/models"
	"encoding/json"
	"log"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	users, err := models.GetUser()
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(users)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
