package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("hi")
	log.Warn("It's a warn log")
	log.Info("It's a info log")
	log.Fatalf("It's a fatal log")
}
