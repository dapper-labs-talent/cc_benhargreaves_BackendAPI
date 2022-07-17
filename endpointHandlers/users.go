package endpointHandlers

import (
	"Dapperlabs_Challenge/models"
	"Dapperlabs_Challenge/utils"
	"encoding/json"
	"log"
	"net/http"
)

//Fetch all users from DB. Requires an valid JWT from authenticated user to complete request
func GetUsers(w http.ResponseWriter, r *http.Request) {
	//Fetch and validate token from header
	token := r.Header.Get("x-authentication-token")
	if token == "" {
		utils.WriteJSONError(w, utils.EmptyJWTErr, http.StatusUnauthorized)
		log.Print(utils.EmptyJWTErr)
		return
	}

	_, err := utils.ValidateToken(token)
	if err != nil {
		utils.WriteJSONError(w, utils.InvalidJWTErr, http.StatusUnauthorized)
		log.Print(err.Error())
		return
	}

	//Token valid - fetch list of all users in DB
	users, err := models.GetUsers()
	if err != nil {
		utils.WriteJSONError(w, utils.GetAllUsersErr, http.StatusInternalServerError)
		log.Print(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string][]models.User{"users": users})
}

//Update the firstname and/or lastname of the currently logged in user. Requires a valid JWT from authenticated user to complete request
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	//Fetch and validate token from header
	token := r.Header.Get("x-authentication-token")
	if token == "" {
		utils.WriteJSONError(w, utils.EmptyJWTErr, http.StatusUnauthorized)
		log.Print(utils.EmptyJWTErr)
		return
	}

	//JWT claims will contain the users email. This will be used to know which user to update
	claims, err := utils.ValidateToken(token)
	if err != nil {
		utils.WriteJSONError(w, utils.InvalidJWTErr, http.StatusUnauthorized)
		log.Print(err.Error())
		return
	}

	var name models.Name
	err = json.NewDecoder(r.Body).Decode(&name)
	if err != nil {
		utils.WriteJSONError(w, utils.MalformedRequestBodyErr, http.StatusBadRequest)
		log.Print(err.Error())
		return
	}

	//Update Firstname and/or lastname of user with email == claims.email
	//First and Last name can both be nil here intentionally. No Validation check is run against the Name struct in this function
	err = models.UpdateUserByEmail(claims.Email, name.FirstName, name.LastName)
	if err != nil {
		utils.WriteJSONError(w, utils.UpdateUserErr, http.StatusInternalServerError)
		log.Print(err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
