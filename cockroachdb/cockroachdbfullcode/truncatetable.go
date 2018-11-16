package cockroachdbfullcode

import (
	"database/sql"
	"fmt"
	"log"
)

func TruncateTable(dbConnection *sql.DB) {
	if _, err := dbConnection.Exec("TRUNCATE TABLE studentdetails"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Truncate table successfully.Now Create table again")
	FeatureSelection(dbConnection)
}
