package Gorm

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Product struct {
	gorm.Model
	Code  int `gorm:"primary_key"`
	Price uint
}

func GORMWithCockroach() {
	const addr = "postgresql://root@localhost:26257/testing?sslmode=disable"
	db, err := gorm.Open("postgres", addr)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.CreateTable(&Product{})

	// Create
	db.Create(&Product{Code: 001, Price: 1000})
}
