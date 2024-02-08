package adapters

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtAdapter struct {
	secret string
}

func NewJwtAdapter() *JwtAdapter {
	return &JwtAdapter{
		secret: os.Getenv("JWT_SECRET"),
	}
}

func (j *JwtAdapter) GenerateToken(key string, value string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims[key] = value
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()

	tokenString, err := token.SignedString([]byte(j.secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *JwtAdapter) ValidateToken(tokenString string) (bool, map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.secret), nil
	})

	if err != nil {
		return false, nil, err
	}

	if tokenDecoded, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true, tokenDecoded, nil
	}

	return false, nil, nil
}
