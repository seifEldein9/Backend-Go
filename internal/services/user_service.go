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
func LoginUser(email, password string) (models.User, bool, error) {
	user, err := repository.GetUserByEmail(email)
	if err != nil {
			return models.User{}, false, err
	}
	isValidPassword := repository.CheckPasswordHash(password, user.Password)
	return user, isValidPassword, nil
}