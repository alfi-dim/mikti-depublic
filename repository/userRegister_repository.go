package repository

import (
	"mikti-depublic/model/domain"

	"gorm.io/gorm"
)

type UserRegisterRepository struct {
	DB *gorm.DB
}

func NewUserRegisterRepository(db *gorm.DB) *UserRegisterRepository {
	return &UserRegisterRepository{DB: db}
}

func (repo *UserRegisterRepository) Save(user domain.User) error {
	result := repo.DB.Create(&user)
	return result.Error
}
