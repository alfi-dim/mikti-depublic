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

func (repo *AdminRepository) Save(admin domain.Admin) (domain.Admin, error) {
	err := repo.DB.Create(&admin).Error
	if err != nil {
		return domain.Admin{}, err
	}
	return admin, nil
}
