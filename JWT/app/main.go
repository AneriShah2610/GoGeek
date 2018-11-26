package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var mySecretKey = []byte("secretKey")

// GenerateToken generate token
func GenerateToken() (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = "Aneri"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	// signature verify
	tokenString, err := token.SignedString(mySecretKey)
	if err != nil {
		log.Fatal("Error while add signature", err)
	}
	return tokenString, err
}
func handleRequest() {
	// Create new mux router
	router := mux.NewRouter()
	router.HandleFunc("/app", homePage).Methods("GET")
	log.Fatal("Error while handling function", http.ListenAndServe(":8000", router))
}
func homePage(writer http.ResponseWriter, request *http.Request) {
	tokenstring, err := GenerateToken()
	if err != nil {
		log.Fatal("Error while generating token", err)
	}
	client := &http.Client{}

	// create a new request to given localhost
	req, _ := http.NewRequest("GET", "http://localhost:9000/server", nil)
	// Append JSON Web Token
	request.Header.Set("Token", tokenstring)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Error while get response from server", err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error while Read body", err)
	}
	fmt.Fprintf(writer, string(body))
}
func main() {
	fmt.Println("My Simple Client ")
	handleRequest()
}
