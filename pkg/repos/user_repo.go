package repository

import (
	"github.com/FKuiv/LocalChat/pkg/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (repo *UserRepo) GetUsers() ([]models.User, error) {
	var users []models.User
	result := repo.db.Find(&users)
	return users, result.Error
}

func (repo *UserRepo) GetExistingUser(username string) *models.User {
	return nil
}

func (repo *UserRepo) CreateUser(user models.User) (*models.User, error) {
	return nil, nil
}
