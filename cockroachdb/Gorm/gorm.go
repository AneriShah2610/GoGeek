package Gorm

import (
	"log"

	// Import GORM-related packages.
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Emp struct {
	empID   int `gorm:"primary_key"`
	empName string
}

func CockroachdbWithGorm() {
	// Connect to the "testing" database as the "root" user.
	const addr = "postgresql://root@localhost:26257/testing?sslmode=disable"
	db, err := gorm.Open("postgres", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Automatically create the "emps" table based on the Emp struct model.
	db.AutoMigrate(&Emp{})
	/*// Insert two rows into the "Emp" table.
	db.Create(&emp{empID: 9, empName: "Pqr"})
	db.Create(&emp{empID: 10, empName: "Lmn"})

	// Print out the balances.
	var emps []emp
	db.Find(&emps)
	fmt.Println("Data of employee table")
	for _, emps := range emps {
		fmt.Println(emps.empID, emps.empName)
	}*/
}
