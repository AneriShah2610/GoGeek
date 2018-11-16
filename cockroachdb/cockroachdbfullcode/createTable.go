package cockroachdbfullcode

import (
	"database/sql"
	"fmt"
	"log"
)

// CreateTable if it not exists
func CreateTable(dbConnection *sql.DB) {
	if _, err := dbConnection.Exec("CREATE TABLE IF NOT EXISTS studentdetails (studentid integer primary key,studentname string NOT NULL,department string,age int)"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table created successfully")
	FeatureSelection(dbConnection)
}
