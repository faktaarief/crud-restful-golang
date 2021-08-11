package services

import (
	"github.com/faktaarief/crud-restful-golang/app/database"
	"github.com/faktaarief/crud-restful-golang/app/models"
)

func CreatePost(objectField *models.Post) (*models.Post, error) {
	user := models.Post{}
	err := database.Connector.Create(objectField).Error

	if err != nil {
		return &models.Post{}, err
	}

	return &user, err
}

func FindAllPosts() (*[]models.Post, error) {
	posts := []models.Post{}
	err := database.Connector.Model(&models.Post{}).Preload("Author").Find(&posts).Error

	if err != nil {
		return &[]models.Post{}, err
	}

	return &posts, err
}

func FindPostById(post_id uint64) (*models.Post, error) {
	post := models.Post{}
	err := database.Connector.Model(&models.Post{}).Preload("Author").First(&post, post_id).Error

	if err != nil {
		return &models.Post{}, err
	}

	return &post, err
}

func DeletePostById(post_id uint64) (*models.Post, error) {
	post := models.Post{}
	err := database.Connector.First(&post, post_id).Delete(&post, post_id).Error

	if err != nil {
		return &models.Post{}, err
	}

	return &post, err
}

func UpdatePostById(objectField *models.Post, post_id uint64) (*models.Post, error) {
	post := models.Post{}
	err := database.Connector.First(&post, post_id).Updates(&objectField).Error

	if err != nil {
		return &models.Post{}, err
	}

	return &post, err
}
