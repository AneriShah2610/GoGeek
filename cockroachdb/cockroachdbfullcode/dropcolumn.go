package cockroachdbfullcode

import (
	"database/sql"
	"fmt"
	"log"
)

func DropColumn(dbConnection *sql.DB) {
	if _, err := dbConnection.Exec("ALTER TABLE studentdetails DROP COLUMN age"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Drop Column successfully")
	FeatureSelection(dbConnection)
}
