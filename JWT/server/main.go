package main

import (
	"fmt"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var mySecretKey = []byte("secretKey")

func homePage(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Secret information")
}

// isAuthorized to check token is authorized or not
func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.Header["Token"] != nil {
			token, err := jwt.Parse(request.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Theer an error")
				}
				return mySecretKey, nil
			})
			if err != nil {
				log.Fatal("error while parsing token : ", err)
			}
			if token.Valid {
				endpoint(writer, request)
			}
		} else {
			fmt.Fprintf(writer, "Not Authorized")
		}
	})
}
func handleRequest() {
	router := mux.NewRouter()
	router.Handle("/server", isAuthorized(homePage))
	log.Fatal("Error at handle", http.ListenAndServe(":9000", router))
}
func main() {
	fmt.Println("My simple server")
	handleRequest()
}
