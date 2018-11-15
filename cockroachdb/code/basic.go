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
	if _, err := db.Exec("insert into testing.emp1 values (5,'Anushka'),(6,'Aayush')"); err != nil {
		log.Fatal(err)
	}
	// Print out data of emp1 table frrom testing database
	rows, err := db.Query("select * from testing.emp1")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	fmt.Println("Details about employee 1 table")
	for rows.Next() {
		var emp_id int
		var emp_name string
		if err := rows.Scan(&emp_id, &emp_name); err != nil {
			log.Fatal(err)
		}
		fmt.Println(emp_id, emp_name)
	}
}
