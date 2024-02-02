package providers

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtProvider struct{}

func (j *JwtProvider) GenerateToken(key string, value string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims[key] = value
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()

	tokenString, err := token.SignedString([]byte("my_secret_key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *JwtProvider) ValidateToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("my_secret_key"), nil
	})

	if err != nil {
		return false, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["user"])
		return true, nil
	}

	return false, nil
}
