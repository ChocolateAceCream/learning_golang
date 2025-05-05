package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type MyClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

var Secret = []byte("secret")

func main() {
	email := "asd@asdf.com"
	token, _ := TokenGenerator(email)
	fmt.Println("Token:", token)
	claims, err := Decoder(token)
	if err != nil {
		fmt.Println("Error decoding token:", err)
		return
	}
	fmt.Println("Decoded Claims:", claims)
	fmt.Println("Email:", claims.Email)
}

func TokenGenerator(email string) (tokenStr string, err error) {
	claims := MyClaims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "me",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(100) * time.Second)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err = token.SignedString(Secret)
	if err != nil {
		fmt.Println("Error signing token:", err)
	}
	return
}

func Decoder(tokenString string) (claimsJSON *MyClaims, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return token., nil
	})
	if err != nil {
		return nil, err
	}

	claimsJSON, ok := token.Claims.(*MyClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return
}
