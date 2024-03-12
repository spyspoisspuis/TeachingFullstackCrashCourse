package repositories

import "backend/users/entities"

type Repository interface {
	InsertUser(userEntity *entities.Users) error
}
