package repository

import (
	"mikti-depublic/model/domain"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (repo *UserRepositoryImpl) FindUserByEmail(email string) (*domain.User, error) {
	tx := repo.db.Begin()
	user := new(domain.User)
	if err := tx.Debug().Where("email = ?", email).Take(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (repo *UserRepositoryImpl) FindAdminByEmail(email string) (*domain.Admin, error) {
	tx := repo.db.Begin()
	admin := new(domain.Admin)
	if err := tx.Debug().Where("email = ?", email).Take(&admin).Error; err != nil {
		return admin, err
	}

	return admin, nil
}
