package database

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
)

type DbConnection struct {
	Dbconn *sql.DB
}

var instance *DbConnection
var once sync.Once

func Connection() *DbConnection {
	connectionString := "postgresql://root@localhost:26257/testing?sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	once.Do(func() {
		fmt.Println("Dtabase coonnection : ", db)
		if err != nil {
			log.Fatal("Error while creating database connection", err)
		}
		fmt.Println("Datbase connected")
		instance = &DbConnection{
			Dbconn: db,
		}
	})
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS college (id integer primary key,name string NOT NULL,address string)"); err != nil {
		log.Fatal("Error at creating table", err)
	}
	fmt.Println("Table created successfully")
	return instance
}
