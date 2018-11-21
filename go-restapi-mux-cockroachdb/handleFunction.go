package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

// AllUsers shows all data of table
func AllUsers(writer http.ResponseWriter, request *http.Request) {
	dbConnection := CreateDbConnection()
	data, err := dbConnection.Query("SELECT *FROM users")
	if err != nil {
		log.Fatal(err)
	}
	var users User
	fmt.Fprintf(writer, `All data in users table`)
	fmt.Fprintf(writer, "\n")
	for data.Next() {
		if err := data.Scan(&users.UserID, &users.UserName, &users.UserEmail, &users.UserContact); err != nil {
			log.Fatal("Error while scanning data", err)
		}
		json.NewEncoder(writer).Encode(users)
	}
}

// NewUser use to create new user
func NewUser(writer http.ResponseWriter, request *http.Request) {
	dbConnection := CreateDbConnection()

	var users User
	_ = json.NewDecoder(request.Body).Decode(&users)

	if _, err := dbConnection.Exec("INSERT INTO users values (4,'Anushka','anushkashah97@gmail.com','857495654')"); err != nil {
		log.Fatal("Error while creating new user", err)
	}
	fmt.Fprintf(writer, `Create new user successfullly`)
}

// DeleteUser use to delete data from table
func DeleteUser(writer http.ResponseWriter, request *http.Request) {
	dbConnection := CreateDbConnection()

	if _, err := dbConnection.Exec("delete from users where userid=4"); err != nil {
		log.Fatal("Error while delete user")
	}
	fmt.Fprintf(writer, `User deleted successfully`)
}
