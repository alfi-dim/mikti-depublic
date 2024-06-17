package service

import (
	"mikti-depublic/helper"
	"mikti-depublic/model/domain"
	"mikti-depublic/repository"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AdminService struct {
	Repo *repository.AdminRepository
	DB   *gorm.DB
}

func NewAdminService(repo *repository.AdminRepository, db *gorm.DB) *AdminService {
	return &AdminService{Repo: repo, DB: db}
}

func (s *AdminService) Register(admin domain.Admin) error {
	adminID, err := helper.GenerateID(s.DB, admin.TableName(), "ADMIN")
	if err != nil {
		return err
	}
	admin.ID = adminID

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	admin.Password = string(hashedPassword)

	err = s.Repo.Save(admin)
	if err != nil {
		return err
	}
	return nil
}
