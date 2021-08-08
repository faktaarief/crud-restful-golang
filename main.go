package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/faktaarief/crud-restful-golang/app/database"
	"github.com/faktaarief/crud-restful-golang/app/models"
	"github.com/faktaarief/crud-restful-golang/app/routes"
	"github.com/joho/godotenv"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic("Error getting env")
	} else {
		fmt.Println("Env value has been loaded")
	}

	initDB()
	initModels()
	routes.Router()
}

func initDB() {
	config :=
		database.Config{
			Host:     os.Getenv("DB_HOST"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASS"),
			DB:       os.Getenv("DB_NAME"),
			Port:     os.Getenv("DB_PORT"),
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(os.Getenv("DB_DRIVER"), connectionString)
	if err != nil {
		panic(err.Error())
	}
}

func initModels() {
	for _, model := range models.RegisterModels() {
		err := database.Connector.AutoMigrate(model.Model).Error
		if err != nil {
			log.Fatalf("Model error %v", err)
		}
	}

	models.AddConstraints()
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(res, req)
	})
}
