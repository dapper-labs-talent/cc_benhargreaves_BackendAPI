package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var secretKey = []byte(os.Getenv("JWTKEY"))

type JWTClaim struct {
	Email string `json:"Email"`
	jwt.RegisteredClaims
}

//Create a new signed JWT with an expiration of +1hour
//Provided email will be included in this JWTs claims
func CreateToken(email string) (string, error) {
	expirationTime := jwt.NewNumericDate(time.Now().Add(1 * time.Hour))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &JWTClaim{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expirationTime,
		},
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

//Validate signed Token, and verifies the token has not expired
//return the tokens Claims
func ValidateToken(tokenString string) (*JWTClaim, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New(UnableToParseJWTClaimsErr)
		return nil, err
	}

	if !claims.VerifyExpiresAt(time.Now().Local(), false) {
		err = errors.New(ExpiredTokenErr)
		return nil, err
	}

	return claims, nil
}
