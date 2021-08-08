package controllers

import (
	"net/http"

	"github.com/faktaarief/crud-restful-golang/app/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateUser(ctx *gin.Context) {
	user := models.User{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	validate := validator.New()

	if err := validate.Struct(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}
