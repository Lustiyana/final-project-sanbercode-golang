package helpers

import (
	"fmt"
	"time"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("1NIP45SW0RD")

type Claims struct {
	ID int64 `json:"id"`
	Email string `json:"email"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(id int64, email string, password string) (string, error) {
	claims := &Claims{
		ID: id,
		Email: email,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
			IssuedAt:  time.Now().Unix(),
			Issuer:    "golang_sanbercode",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("Invalid token")
	}

	return claims, nil
}


func ExtractToken(auth string) (string, error) {
	token := strings.TrimPrefix(auth, "Bearer ")
	if token == auth {
		return "", http.ErrNotSupported
	}
	return token, nil
}