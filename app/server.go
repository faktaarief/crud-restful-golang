package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/faktaarief/crud-restful-golang/app/models"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(dbDriver, dbUser, dbPass, dbPort, dbHost, dbName string) {
	var err error

	if dbDriver == "mysql" {
		dbURL := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			dbUser, dbPass, dbHost, dbPort, dbName,
		)

		server.DB, err = gorm.Open(dbDriver, dbURL)

		if err != nil {
			fmt.Printf("Cannot connect to %s database", dbDriver)
			log.Fatal("Message: ", err)
		} else {
			fmt.Printf("Successfully connected to %s database", dbDriver)
		}
	}

	server.DB.Debug().AutoMigrate(
		&models.User{},
		&models.Post{},
	)

	err = server.DB.Debug().Model(&models.Post{}).AddForeignKey(
		"author_id",
		"users(id)",
		"cascade",
		"cascade",
	).Error

	if err != nil {
		log.Fatalf("Attaching ForeignKey Error: %v", err)
	}

	server.Router = mux.NewRouter()
}

func (server *Server) Run(PORT string) {
	fmt.Printf("Listening to port %s", PORT)
	log.Fatal(http.ListenAndServe(PORT, server.Router))
}
