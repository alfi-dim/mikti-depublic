package service

import (
	"mikti-depublic/helper"
	"mikti-depublic/model/domain"
	"mikti-depublic/model/web"
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

func (s *AdminService) Register(request web.AdminRegisterRequest) (domain.Admin, error) {
	admin := domain.Admin{}
	adminID, err := helper.GenerateID(s.DB, admin.TableName(), "ADMIN")
	if err != nil {
		return domain.Admin{}, nil
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		return domain.Admin{}, nil
	}
	adminReq := domain.Admin {
		ID: adminID,
		Name: request.Name,
		Username: request.Username,
		Email: request.Email,
		Password: string(hashedPassword),
	}

	newAdmin, err := s.Repo.Save(adminReq)
	if err != nil {
		return domain.Admin{}, err
	}
	return newAdmin, nil
}
