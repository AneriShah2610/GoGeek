package main

import (
	"fmt"

	"github.com/AneriShah2610/GoGeek/GoModule/mo"

	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("hi")
	log.Warn("It's a warn log")
	log.Info("It's a info log")

	mo.Mo()
	log.Fatalf("It's a fatal log")
}
