package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
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

	for _, model := range RegisterModels() {
		err = server.DB.Debug().AutoMigrate(model.Model).Error

		if err != nil {
			log.Fatalf("Model error %v", err)
		}
	}

	server.AddConstraints()
	server.Router = mux.NewRouter()
	server.Router.Use(commonMiddleware)
	server.Routes()
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(res, req)
	})
}

func (server *Server) Run(PORT string) {
	fmt.Printf("Listening to port %s", PORT)
	log.Fatal(http.ListenAndServe(PORT, server.Router))
}
