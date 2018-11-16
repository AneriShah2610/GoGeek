package cockroachdbfullcode

import (
	"database/sql"
	"fmt"
	"log"
)

func DropTable(dbConnection *sql.DB) {
	if _, err := dbConnection.Exec("DROP TABLE studentdetails"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Drop Table successfully")
	FeatureSelection(dbConnection)
}
