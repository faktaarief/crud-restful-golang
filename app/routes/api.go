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
	router.GET("/users", controllers.FindAllUsers)
	router.GET("/users/:id", controllers.FindUserById)
	router.POST("/users/create", controllers.CreateUser)
	router.DELETE("/users/:id", controllers.DeleteUserById)
	router.PUT("/users/:id", controllers.UpdateUserById)

	router.Run(":8080")
	log.Println("Starting the HTTP Server on Port", PORT)
}
