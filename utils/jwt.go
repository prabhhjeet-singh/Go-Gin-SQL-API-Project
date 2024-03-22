package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const key = "nuclearcodes"

func GenerateToken(email string, userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(key))
}

func VerifyToken(tokenString string) (error, int) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	_, ok := token.Method.(*jwt.SigningMethodHMAC)

	if !ok {
		return errors.New("invalid signing method"), 0
	}

	if err != nil {
		return errors.New("could not parse token"), 0
	}

	if !token.Valid {
		return errors.New("invalid token"), 0
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return errors.New("invalid token claims"), 0
	}

	userId := int(claims["userId"].(float64))

	return nil, userId
}
