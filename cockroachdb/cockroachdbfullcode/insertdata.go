package cockroachdbfullcode

import (
	"database/sql"
	"fmt"
	"log"
)

func InsertData(dbConnection *sql.DB) {
	if _, err := dbConnection.Exec("INSERT INTO studentdetails VALUES (1,'Aneri','Computer',22),(2,'Anokhi','Pharmacy',22),(3,'Aayush','ICT',20),(4,'Anushka','Dental',21)"); err != nil {
		log.Fatal("Error while inserting data into studentdetails table", err)
	}
	fmt.Println("Data inserted successfully")
	FeatureSelection(dbConnection)
}
