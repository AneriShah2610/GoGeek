package main

import (
	"github.com/AneriShah2610/GoGeek/cockroachdb/Gorm"
	"github.com/AneriShah2610/GoGeek/cockroachdb/cockroachdbfullcode"
	"github.com/AneriShah2610/GoGeek/cockroachdb/code"
	_ "github.com/lib/pq"
)

// Code,Gorm are subpackgaes of cocckroachdb directory

func main() {
	// code contains simple cockroachdb example with go
	code.BasicCockroachdbCode()
	// code is done using GORM+Go+Cockroachdb and db.AutoMigrate method to auto create table
	Gorm.CockroachdbWithGorm()
	// code is done using GORM+Go+Cockroachdb and db.CreateTable method to create table
	Gorm.GORMWithCockroach()
	// this code has different function for different functionalities
	code.Cockroachdb()
	// code contains the features like Go+Cockroachdb+postgresql_driver+function+struct
	code.CockroachDbwithStruct()
	// code contains the features like Go+Cockroachdb+postgresql_driver+function with update function
	code.UpdateDataFunc()
	//All cockroachdb operation
	cockroachdbfullcode.DbConnectionCreate()
}
