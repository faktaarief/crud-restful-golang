package routes

import (
	"fmt"
	"log"
	"os"

	"github.com/faktaarief/crud-restful-golang/app/controllers"
	"github.com/gin-gonic/gin"
)

func Router() {
	var PORT = fmt.Sprintf(":%s", os.Getenv("PORT"))
	router := gin.Default()

	router.GET("/", controllers.Home)
	router.POST("/users/create", controllers.CreateUser)

	router.Run(":8080")
	log.Println("Starting the HTTP Server on Port", PORT)
}
