package repository

import (
	"mikti-depublic/model/domain"

	"gorm.io/gorm"
)

type AdminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *AdminRepository {
	return &AdminRepository{DB: db}
}

func (repo *AdminRepository) Save(admin domain.Admin) error {
	result := repo.DB.Create(&admin)
	return result.Error
}
