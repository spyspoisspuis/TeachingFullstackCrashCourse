package repositories

import (
	"backend/users/entities"

	"gorm.io/gorm"
)

type postgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) Repository {
	return &postgresRepository{db: db}
}

func (p *postgresRepository) InsertUser(userEntity *entities.Users) error {
	result := p.db.Create(userEntity)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (p *postgresRepository) GetAllUser() ([]entities.Users, error) {
	var users []entities.Users
	result := p.db.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}