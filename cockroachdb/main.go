package main

import (
	"github.com/AneriShah2610/GoGeek/cockroachdb/Gorm"
	"github.com/AneriShah2610/GoGeek/cockroachdb/code"
	_ "github.com/lib/pq"
)

func main() {
	code.BasicCockroachdbCode()
	//Gorm.CockroachdbWithGorm()
	Gorm.GORMWithCockroach()
}
