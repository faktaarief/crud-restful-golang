package controllers

import (
	"net/http"

	"github.com/faktaarief/crud-restful-golang/app/services"
	"github.com/gin-gonic/gin"
)

func FindAllPosts(ctx *gin.Context) {
	posts, err := services.FindAllPosts()

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"datas": posts,
	})
}
