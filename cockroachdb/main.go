package main

import (
	"github.com/AneriShah2610/GoGeek/cockroachdb/Gorm"
	"github.com/AneriShah2610/GoGeek/cockroachdb/code"
	_ "github.com/lib/pq"
)

func main() {
	// Code is a subpackga eof cocckroach db and it has a simple cockroachdb example with go
	code.BasicCockroachdbCode()
	//Gorm is a subpackga eof cocckroach db and it use GORM+Go+Cockroachdb and db.AutoMigrate method to auto create table
	Gorm.CockroachdbWithGorm()
	//Gorm is a subpackga eof cocckroach db and it use GORM+Go+Cockroachdb and db.CreateTable method to create table
	Gorm.GORMWithCockroach()
}
