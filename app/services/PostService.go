package services

import (
	"github.com/faktaarief/crud-restful-golang/app/database"
	"github.com/faktaarief/crud-restful-golang/app/models"
)

func FindAllPosts() (*[]models.Post, error) {
	posts := []models.Post{}
	err := database.Connector.Model(&models.Post{}).Preload("Author").Find(&posts).Error

	if err != nil {
		return &[]models.Post{}, err
	}

	return &posts, err
}
