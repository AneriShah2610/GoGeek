package cockroachdbfullcode

import (
	"database/sql"
	"fmt"
	"log"
)

func DbConnectionCreate() {
	connectionString := "postgresql://root@localhost:26257/testing?sslmode=disable"
	dbConnection, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("Error while connecting database :", err)
	}
	fmt.Println("Database connection : ", dbConnection)
	FeatureSelection(dbConnection)
}
func FeatureSelection(dbConnection *sql.DB) {
	var operation string
	fmt.Println("Select any one from below list : \n 1. Create Table \n 2. Insert Data \n 3. Update data \n 4. Delete data \n 5. Drop column \n 6. Drop table \n 7. Truncate table \n 8. Print data")
	fmt.Scanln(&operation)
	switch {
	case operation == "Create Table" || operation == "1":
		CreateTable(dbConnection)
	case operation == "Insert Data" || operation == "2":
		InsertData(dbConnection)
	case operation == "Update data" || operation == "3":
		UpdateData(dbConnection)
	case operation == "Delete data" || operation == "4":
		DeleteData(dbConnection)
	case operation == "Drop column" || operation == "5":
		DropColumn(dbConnection)
	case operation == "Drop table" || operation == "6":
		DropTable(dbConnection)
	case operation == "Truncate table" || operation == "7":
		TruncateTable(dbConnection)
	case operation == "Print data" || operation == "8":
		PrintTableData(dbConnection)
	}
}
