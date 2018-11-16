package cockroachdbfullcode

import (
	"database/sql"
	"fmt"
	"log"
)

func UpdateData(dbConnection *sql.DB) {
	if _, err := dbConnection.Exec("UPDATE studentdetails SET studentname='Ani' where studentid=1"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Update data successfully")
	FeatureSelection(dbConnection)
}
