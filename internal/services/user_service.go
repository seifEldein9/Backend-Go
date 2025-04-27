package services

import (
	"go-backend-app/internal/models"
	"go-backend-app/internal/repository"
)

func GetUsers() ([]models.User, error) {
    return repository.GetAllUsers()
}

func AddUser(user models.User) error {
    return repository.CreateUser(user)
}
