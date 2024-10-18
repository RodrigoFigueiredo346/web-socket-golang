package services

import (
	"fmt"
	"main/internal/config"
	"net/http"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

func CreateToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userID": userID,
			"exp":    time.Now().Add(time.Minute * 1000).Unix(),
		})
	secretKey := config.GetConfig().SecretKey

	return token.SignedString(secretKey)
}

func VerifyToken(r *http.Request) error {
	secretKey := config.GetConfig().SecretKey

	tokenInHeader := r.Header.Get("Authorization")
	tokenParts := strings.Split(tokenInHeader, " ")
	if len(tokenParts) != 2 {
		return fmt.Errorf("token invalid")
	}
	tokenString := tokenParts[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

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

func IsTokenExpired(token *jwt.Token) bool {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return true
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		return true
	}

	return time.Now().Unix() > int64(exp)
}

func ExtractUserIDFromToken(tokenString string) (string, error) {
	secretKey := config.GetConfig().SecretKey

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return "", err
	}

	// Extraindo as claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if userID, ok := claims["userID"].(string); ok {
			return userID, nil
		}
		return "", fmt.Errorf("userID not found in token")
	} else {
		return "", fmt.Errorf("invalid token")
	}
}
