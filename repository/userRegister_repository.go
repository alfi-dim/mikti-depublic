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

func (repo *UserRegisterRepository) Save(user domain.User) (domain.User, error) {
	err := repo.DB.Create(&user).Error

	if err != nil {
		return domain.User{}, err
	}

	return user, nil
	
}
