package main

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//var mySigningKey = os.Get("MY_JWT_TOKEN")

var mySigningKey = []byte("superSecretPhrase")

func generateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("something went wrong", err.Error())
		return "", err
	}
	return tokenString, nil
}

func main() {
	fmt.Println("My simple client")

	tokenString, err := generateJWT()

	if err != nil {
		fmt.Println("error generating tokenString")
	}

	fmt.Println(tokenString)
}
