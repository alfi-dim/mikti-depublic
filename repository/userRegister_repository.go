package repository

import (
	"mikti-depublic/model/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (repo *UserRepository) Save(user domain.User) error {
	result := repo.DB.Create(&user)
	return result.Error
}
