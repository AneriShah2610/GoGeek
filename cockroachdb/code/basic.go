package code

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func BasicCockroachdbCode() {
	// Connect to the database
	db, err := sql.Open("postgres", "postgresql://root@localhost:26257/testing?sslmode=disable")
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	// Insert into table
	/*if _, err := db.Exec("insert into testing.emp1 values (7,'Abc'),(8,'Xyz')"); err != nil {
		log.Fatal(err)
	}*/

	// Print out data of emp1 table frrom testing database
	rows, err := db.Query("select * from testing.emp1")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	fmt.Println("Details about employee 1 table")
	for rows.Next() {
		var empid int      // column1 in emp1 table
		var empname string //column2 in emp1 table
		if err := rows.Scan(&empid, &empname); err != nil {
			log.Fatal(err)
		}
		fmt.Println(empid, empname)
	}
}
