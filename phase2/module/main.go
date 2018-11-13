package main

import (
	"github.com/kr/pretty"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("This is a log info")
	log.Warn("This is log warning")
	pretty.Println("Hello world")
	log.Fatal("This is lof fatal")
}
