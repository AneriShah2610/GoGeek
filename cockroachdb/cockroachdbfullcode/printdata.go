package cockroachdbfullcode

import (
	"database/sql"
	"fmt"
	"log"
)

type Student struct {
	studentid   int
	studentname string
	department  string
	age         int
}

func PrintTableData(dbConnection *sql.DB) {
	data, err := dbConnection.Query("SELECT *FROM studentdetails")
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	student := Student{}
	fmt.Println("Print details of studentdetails table:")
	for data.Next() {
		if err := data.Scan(&student.studentid, &student.studentname, &student.department, &student.age); err != nil {
			log.Fatal(err)
		}
		fmt.Println(student)
	}
	FeatureSelection(dbConnection)
}
