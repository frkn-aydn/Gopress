package jsonWebToken

import (
	jwt "github.com/dgrijalva/jwt-go"
)

func Make() (string, error) {
	type User struct {
		jwt.StandardClaims
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &User{
		Name: "otiai10",
		Age:  30,
	})
	hmacSampleSecret := []byte("t1LrbFpKG5v4ENntnFeUWHOAY7cYHLUS")
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSampleSecret)

	return tokenString, err
}
