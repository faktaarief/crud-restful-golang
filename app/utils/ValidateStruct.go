package utils

import (
	"github.com/faktaarief/crud-restful-golang/app/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidateStructPost(post *models.Post, ctx *gin.Context) error {
	if err := ctx.ShouldBindJSON(&post); err != nil {
		return err
	}

	validate := validator.New()

	if err := validate.Struct(post); err != nil {
		return err
	}

	return nil
}
