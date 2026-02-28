package main

import (
	"io"
	"log"
	"os"

	"github.com/bulat-abd/go1fl-final/pkg/db"
	"github.com/bulat-abd/go1fl-final/pkg/server"

	"github.com/sirupsen/logrus"
)

func initLogging() {
	file, err := os.OpenFile("application.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	mw := io.MultiWriter(os.Stdout, file)
	logrus.SetOutput(mw)
	logrus.SetLevel(logrus.WarnLevel)
}

func main() {
	initLogging()
	database, _ := db.Init()
	defer database.Close()
	server.Run(database)
}
