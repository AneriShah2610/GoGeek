package cockroachdbfullcode

import (
	"database/sql"
	"fmt"
	"log"
)

func DeleteData(dbConnection *sql.DB) {
	var studentid int
	fmt.Println("Enter student id which u want to delete data : ")
	fmt.Scanln(&studentid)
	if _, err := dbConnection.Exec("DELETE from studentdetails WHERE studentid=studentid"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Delete data successfully")
	FeatureSelection(dbConnection)
}
