package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var dbConnection *sql.DB

// User model
type User struct {
	UserID      int
	UserName    string
	UserEmail   string
	UserContact string
}

// CreateDbConnection create database connection
func CreateDbConnection() *sql.DB {
	connectionString := "postgresql://root@localhost:26257/testing?sslmode=disable" // db driver
	dbConnection, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("Error while creating database connection", err)
	}

	return dbConnection
}

// CreateTable create table if not exists
func CreateTable(dbConnection *sql.DB) {
	if _, err := dbConnection.Exec("CREATE TABLE IF NOT EXISTS users (userid integer primary key,username string NOT NULL,email string NOT NULL,contactno string)"); err != nil {
		log.Fatal("Error at creating users table", err)
	}
}
