package code

import (
	"database/sql"
	"fmt"
	"log"
)

type (
	Recipe struct {
		RecipeID    int
		RecipeName  string
		Ingrediants string
	}
)

func CockroachDbwithStruct() {
	createdbConnection()

}
func createdbConnection() {
	connectionString := "postgresql://root@localhost:26257/testing?sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(db)
	// create table recipes if not exists
	createTableRecipe(db)
}

func createTableRecipe(db *sql.DB) {
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS recipes(RecipeId integer primary key,RecipeName string NOT NULL,ingrediants string)"); err != nil {
		log.Fatal("Error while creating table", err)
	}
	readTableData(db)
}

func readTableData(db *sql.DB) {
	data, err := db.Query("select *from recipes")
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	recipe := Recipe{}
	fmt.Println("Data of recipes table : ")
	for data.Next() {
		if err := data.Scan(&recipe.RecipeID, &recipe.RecipeName, &recipe.Ingrediants); err != nil {
			log.Fatal(err)
		}
		fmt.Println(recipe)
	}
}
