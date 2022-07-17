package endpointHandlers

import (
	"Dapperlabs_Challenge/models"
	"Dapperlabs_Challenge/utils"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

//Create new user in Postgres, and return a JWT token to be used to authenticate future requests for this user
//Expects Request body with the following fields:
//	{
// 		"email": "test@axiomzen.co",
// 		"password": "axiomzen",
// 		"firstName": "Alex",
// 		"lastName": "Zimmerman"
//	}
func Signup(w http.ResponseWriter, r *http.Request) {
	//Decode and marshal request Json into user struct
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.WriteJSONError(w, utils.MalformedRequestBodyErr, http.StatusBadRequest)
		log.Print(err.Error())
		return
	}

	//Validate required fields present for new user
	err = validator.New().Struct(user)
	if err != nil {
		utils.WriteJSONError(w, err.Error(), http.StatusBadRequest)
		log.Print(err.Error())
		return
	}

	//Ensure user doesnt already exist
	_, err = models.GetUserByEmail(user.Email)
	//A nil err would suggest that a row was returned in the DB when looking up the user
	if err == nil {
		utils.WriteJSONError(w, utils.UserAlreadyExistsErr, http.StatusBadRequest)
		return
	}
	//ErrNoRows is what we expect to be returned - if its anything else, something went wrong
	if err != sql.ErrNoRows {
		utils.WriteJSONError(w, utils.UnknownErr, http.StatusInternalServerError)
		log.Print(err.Error())
		return
	}

	//Green light - time to create new user and generate a token
	err = models.CreateUser(user)
	if err != nil {
		utils.WriteJSONError(w, utils.InsertNewUserErr, http.StatusInternalServerError)
		log.Print(err.Error())
		return
	}
	token, err := utils.CreateToken(user.Email)
	if err != nil {
		utils.WriteJSONError(w, utils.CreateNewJWTErr, http.StatusInternalServerError)
		log.Print(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"token": token})

}

//Validate email and password for existing user - and returns a valid JWT token
//Expects Request body with the following fields:
//	{
// 		"email": "test@axiomzen.co",
// 		"password": "axiomzen"
//	}
func Login(w http.ResponseWriter, r *http.Request) {
	//Decode and marshal request body into credentials struct
	var creds models.Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		utils.WriteJSONError(w, utils.MalformedRequestBodyErr, http.StatusBadRequest)
		log.Print(err.Error())
		return
	}

	//make sure required fields present for login
	err = validator.New().Struct(creds)
	if err != nil {
		utils.WriteJSONError(w, err.Error(), http.StatusBadRequest)
		log.Print(err.Error())
		return
	}

	//fetch user from DB. Will return error if no user by that email exists
	user, err := models.GetUserByEmail(creds.Email)
	if err != nil {
		utils.WriteJSONError(w, utils.GetUserByEmailErr, http.StatusNotFound)
		log.Print(err.Error())
		return
	}

	//Compare provided password with hash from DB
	err = utils.ComparePasswordHash(creds.Password, user.Password)
	if err != nil {
		utils.WriteJSONError(w, utils.InvalidPasswordErr, http.StatusUnauthorized)
		log.Print(err.Error())
		return
	}

	//Password verified - generate a new token for this user
	token, err := utils.CreateToken(user.Email)
	if err != nil {
		utils.WriteJSONError(w, utils.CreateNewJWTErr, http.StatusInternalServerError)
		log.Print(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
