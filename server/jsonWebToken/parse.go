package jsonWebToken

import (
	jwt "github.com/dgrijalva/jwt-go"
)

func ParseToken(unparsedToken string) (interface{}, error) {
	token, err := jwt.Parse(unparsedToken, func(token *jwt.Token) (interface{}, error) {
		return []byte("t1LrbFpKG5v4ENntnFeUWHOAY7cYHLUS"), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
