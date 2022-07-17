package models

import (
	"Dapperlabs_Challenge/utils"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	Credentials
	Name
}

type Credentials struct {
	Email    string `json:"email" bun:",notnull,unique" validate:"required,email"`
	Password string `json:"password,omitempty" bun:",notnull" validate:"required"`
}

type Name struct {
	FirstName *string `json:"firstName,omitempty" validate:"required"`
	LastName  *string `json:"lastName,omitempty" validate:"required"`
}

//Returns the email, firstname, and lastname of all users currently in the DB
func GetUsers() ([]User, error) {
	rows, err := DB.Query("SELECT email, firstName, lastName FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var u User

		err := rows.Scan(&u.Email, &u.FirstName, &u.LastName)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

//Fetch the first user that matches the provided email.
//Returns sql.ErrNoRows if no matches were found
func GetUserByEmail(email string) (User, error) {
	var u User
	err := DB.QueryRow("SELECT email, password, firstName, lastName FROM users WHERE email = ?", email).Scan(&u.Email, &u.Password, &u.FirstName, &u.LastName)
	if err != nil {
		return u, err
	}
	return u, nil
}

//Update the firstname and/or lastname of the user matching the provided email
//firstName and lastName can be nil to keep the values as they currently stand in the DB
func UpdateUserByEmail(email string, firstName *string, lastName *string) error {
	// If both fields were nil, no need to make DB call. Just return as success
	if firstName == nil && lastName == nil {
		return nil
	}

	//Update names in DB if provided values arent nil. Otherwise keep them as they currently are in DB
	sqlStatement := `
		UPDATE users 
		SET   
			firstName = COALESCE(?, firstName),
			lastName = COALESCE(?, lastName) 
		WHERE email = ?
	`
	_, err := DB.Exec(sqlStatement, firstName, lastName, email)
	return err
}

//Create new user in the DB
//user.Password will be hashed before insert
func CreateUser(user User) error {
	var err error
	user.Password, err = utils.GenerateHashedPassword(user.Password)
	if err != nil {
		return err
	}
	sqlStatement := `
		INSERT INTO users (email, password, firstname, lastname)
		VALUES (?, ?, ?, ?)`
	_, err = DB.Exec(sqlStatement, user.Email, user.Password, user.FirstName, user.LastName)
	if err != nil {
		return err
	}

	return nil
}
