package code

import (
	"database/sql"
	"fmt"
	"log"
)

// UpdateDataFunc to update data of recipe table
func UpdateDataFunc() {
	dbCheckConnection()
}
func dbCheckConnection() {
	connString := "postgresql://root@localhost:26257/testing?sslmode=disable"
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
	updateTableData(db)
}
func updateTableData(db *sql.DB) {
	if _, err := db.Exec("UPDATE recipes SET recipename='filtercoffee', ingrediants='dark coffee,chocolate powder,milk,water,flavour,chocolate sauce,sugar' WHERE recipeid=4"); err != nil {
		log.Fatal(err)
	}
	printModifiedData(db)
}
func printModifiedData(db *sql.DB) {
	data, err := db.Query("select *from recipes")
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	fmt.Println("Data of recipes table : ")
	for data.Next() {
		var recipeid int
		var recipename, ingrediants string
		if err := data.Scan(&recipeid, &recipename, &ingrediants); err != nil {
			log.Fatal(err)
		}
		fmt.Println(recipeid, recipename, ingrediants)
	}
}
