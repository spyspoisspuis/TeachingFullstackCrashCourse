package usecases

import (
	"backend/users/entities"
	"backend/users/helper"
	"backend/users/models"
	"backend/users/repositories"
)

type UsecaseImpl struct {
	Repository repositories.Repository
}

func NewUsecaseImpl(repository repositories.Repository) Usecase {
	return &UsecaseImpl{
		Repository: repository,
	}
}

func (u *UsecaseImpl) InsertUser(userModel *models.Users) error {
	// models , username , password
	uuid, err := helper.GenerateUUID()
	if err != nil {
		return err
	}

	hashedPassword, err := helper.GenerateHashedPassword(userModel.Password)
	if err != nil {
		return err
	}

	// populate entity field
	userEntity := &entities.Users{
		Id:       uuid,
		Username: userModel.Username,
		Password: hashedPassword,
	}

	// insert user to db
	err = u.Repository.InsertUser(userEntity)
	if err != nil {
		return err
	}

	return nil
}
