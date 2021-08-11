package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/faktaarief/crud-restful-golang/app/models"
	"github.com/faktaarief/crud-restful-golang/app/services"
	"github.com/faktaarief/crud-restful-golang/app/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateUser(ctx *gin.Context) {
	user := models.User{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	validate := validator.New()

	if err := validate.Struct(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := utils.BeforeSaveUser(&user)
	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = services.CreateUser(&user)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func FindAllUsers(ctx *gin.Context) {
	users, err := services.FindAllUsers()

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Users not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"datas": users,
	})
}

func FindUserById(ctx *gin.Context) {
	user_id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	user, err := services.FindUserById(uint32(user_id))

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func DeleteUserById(ctx *gin.Context) {
	user_id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	_, err := services.DeleteUserById(uint32(user_id))

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User Successfully Deleted!",
	})
}

func UpdateUserById(ctx *gin.Context) {
	user := models.User{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	validate := validator.New()

	if err := validate.Struct(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := utils.BeforeSaveUser(&user)
	if err != nil {
		log.Fatal(err)
		return
	}

	user_id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	_, err = services.UpdateUserById(&user, uint32(user_id))

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
