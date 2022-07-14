package models

import (
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID        int64  `bun:,pk,autoincrement`
	Email     string `bun:,notnull,unique`
	Password  string `bun:,notnull`
	FirstName string
	LastName  string
}

func GetUser() ([]User, error) {
	rows, err := DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var u User

		err := rows.Scan(&u.ID, &u.Email, &u.Password, &u.FirstName, &u.LastName)
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
