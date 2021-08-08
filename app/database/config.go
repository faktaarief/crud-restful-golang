package database

import "fmt"

type Config struct {
	Host     string
	User     string
	Password string
	DB       string
	Port     string
}

var GetConnectionString = func(config Config) string {
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.User, config.Password, config.Host, config.Port, config.DB,
	)

	return connectionString
}
