package main

import (
	"fmt"
	"log"
	"os"

	"github.com/faktaarief/crud-restful-golang/app"
	"github.com/joho/godotenv"
)

var server = app.Server{}

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	dbDriver := os.Getenv("DB_DRIVER")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	server.Initialize(dbDriver, dbUser, dbPass, dbPort, dbHost, dbName)
	server.Run(":4001")
}
