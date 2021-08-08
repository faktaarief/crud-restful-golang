package database

import (
	"log"

	"github.com/jinzhu/gorm"
)

var Connector *gorm.DB

func Connect(dbDriver, connectionString string) error {
	var err error
	Connector, err = gorm.Open(dbDriver, connectionString)
	if err != nil {
		return err
	}
	log.Println("Connection was Successfully!")
	return nil
}
