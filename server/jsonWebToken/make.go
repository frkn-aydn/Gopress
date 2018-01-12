package jsonWebToken

import (
	jwt "github.com/dgrijalva/jwt-go"
)

func Make(tokenSource jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenSource)

	hmacSampleSecret := []byte("t1LrbFpKG5v4ENntnFeUWHOAY7cYHLUS")
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSampleSecret)

	return tokenString, err
}
