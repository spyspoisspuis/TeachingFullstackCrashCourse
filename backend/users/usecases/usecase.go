package usecases

import "backend/users/models"

type Usecase interface {
	InsertUser(userModel *models.Users) error
	GetAllUser() ([]models.DisplayUser, error)
}
