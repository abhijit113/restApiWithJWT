package main

import (
	"fmt"
	"log"
	"net/http"
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

func homePage(w http.ResponseWriter, r *http.Request) {
	validToken, err := generateJWT()

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprintf(w, validToken)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":11000", nil))
}

func main() {
	handleRequests()
}
