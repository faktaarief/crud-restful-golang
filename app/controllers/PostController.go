package controllers

import (
	"net/http"
	"strconv"

	"github.com/faktaarief/crud-restful-golang/app/models"
	"github.com/faktaarief/crud-restful-golang/app/services"
	"github.com/faktaarief/crud-restful-golang/app/utils"
	"github.com/gin-gonic/gin"
)

func CreatePost(ctx *gin.Context) {
	post := models.Post{}
	err := utils.ValidateStructPost(&post, ctx)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err = services.CreatePost(&post)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}

	result, err := services.FindPostById(post.ID)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

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

func FindPostById(ctx *gin.Context) {
	post_id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	post, err := services.FindPostById(uint64(post_id))

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": post,
	})
}

func DeletePostById(ctx *gin.Context) {
	post_id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	_, err := services.DeletePostById(uint64(post_id))

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Post Successfully Deleted!",
	})
}

func UpdatePostById(ctx *gin.Context) {
	post := models.Post{}
	err := utils.ValidateStructPost(&post, ctx)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	post_id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	_, err = services.UpdatePostById(&post, uint64(post_id))

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err,
		})
		return
	}

	result, err := services.FindPostById(post_id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
