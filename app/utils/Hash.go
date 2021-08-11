package utils

import (
	"github.com/faktaarief/crud-restful-golang/app/models"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func BeforeSave(objectField *models.User) error {
	hashedPassword, err := HashPassword(objectField.Password)

	if err != nil {
		return err
	}

	objectField.Password = string(hashedPassword)
	return nil
}
