package code

import (
	"database/sql"
	"fmt"
	"log"
)

func Cockroachdb() {
	// create database connection
	db := dbConnection()
	// Create tabel
	createTable(db)
	//function to insert data in account table
	insertData(db)
}

// function connect the database
func dbConnection() *sql.DB {
	connectString := "postgresql://root@localhost:26257/testing?sslmode=disable"
	db, err := sql.Open("postgres", connectString)
	if err != nil {
		log.Fatal("Error while connecting database", err)
	}
	return db
}

// function to create table if it not existes in database
func createTable(db *sql.DB) {
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS accounts(id integer primary key,balance integer)"); err != nil {
		log.Fatal("Error while creating table", err)
	}
}

// insert data into table
func insertData(db *sql.DB) {
	if _, err := db.Exec("INSERT INTO accounts VALUES (103,3000),(104,4000)"); err != nil {
		log.Fatal(err)
	}
	printData(db)
}

// print data of table
func printData(db *sql.DB) {
	rows, err := db.Query("SELECT *FROM accounts")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	fmt.Println("Acoount table details:")
	for rows.Next() {
		var id, balance int
		if err := rows.Scan(&id, &balance); err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, balance)
	}
}
