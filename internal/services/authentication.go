package services

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

func CreateToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userID": userID,
			"exp":    time.Now().Add(time.Minute * 10).Unix(),
		})

	secretKey := GetConfig().SecretKey

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	secretKey := GetConfig().SecretKey

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
