package server

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secret_key")
var refreshKey = []byte("refresh_secret_key")

func CreateToken(userID uint64) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Minute * 24).Unix()

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenStr string) (bool, error) {
	claims := &jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return false, err
	}

	if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
		_ = (*claims)["user_id"]
		return true, nil
	}

	return false, nil
}

func RefreshToken(refreshTokenStr string, userID uint64) (string, error) {
	claims := &jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(refreshTokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return refreshKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
		userIdFromToken := uint64((*claims)["user_id"].(float64))
		if userIdFromToken != userID {
			return "", fmt.Errorf("invalid refresh token")
		}

		newToken, err := CreateToken(userIdFromToken)
		if err != nil {
			return "", err
		}

		return newToken, nil
	}

	return "", fmt.Errorf("invalid refresh token")
}

func CreateRefreshToken(userID uint64) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(refreshKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
