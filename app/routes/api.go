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

	home := router.Group("/")
	{
		home.GET("/", controllers.Home)
	}

	users := router.Group("/users")
	{
		users.GET("/", controllers.FindAllUsers)
		users.GET("/:id", controllers.FindUserById)
		users.POST("/create", controllers.CreateUser)
		users.DELETE("/:id", controllers.DeleteUserById)
		users.PUT("/:id", controllers.UpdateUserById)
	}

	posts := router.Group("/posts")
	{
		posts.GET("/", controllers.FindAllPosts)
		posts.GET("/:id", controllers.FindPostById)
		posts.POST("/", controllers.CreatePost)
		posts.DELETE("/:id", controllers.DeletePostById)
		posts.PUT("/:id", controllers.UpdatePostById)
	}

	router.Run(":8080")
	log.Println("Starting the HTTP Server on Port", PORT)
}
