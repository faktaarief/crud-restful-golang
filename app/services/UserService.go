package services

import (
	"github.com/faktaarief/crud-restful-golang/app/database"
	"github.com/faktaarief/crud-restful-golang/app/models"
)

func CreateUser(objectField *models.User) (*models.User, error) {
	user := models.User{}
	err := database.Connector.Create(objectField).Error

	if err != nil {
		return &models.User{}, err
	}

	return &user, err
}

func FindAllUsers() (*[]models.User, error) {
	users := []models.User{}
	err := database.Connector.Find(&users).Error

	if err != nil {
		return &[]models.User{}, err
	}

	return &users, err
}

func FindUserById(user_id uint32) (*models.User, error) {
	user := models.User{}

	err := database.Connector.First(&user, user_id).Error
	if err != nil {
		return &models.User{}, err
	}

	return &user, err
}

func DeleteUserById(user_id uint32) (*models.User, error) {
	user := models.User{}
	err := database.Connector.First(&user, user_id).Delete(&user, user_id).Error

	if err != nil {
		return &models.User{}, err
	}

	return &user, err
}

func UpdateUserById(objectField *models.User, user_id uint32) (*models.User, error) {
	user := models.User{}
	err := database.Connector.First(&user, user_id).Updates(&objectField).Error

	if err != nil {
		return &models.User{}, err
	}

	return &user, err
}
