package service

import (
	"mikti-depublic/model/domain"
	"mikti-depublic/repository"

	"golang.org/x/crypto/bcrypt"
)

type AdminService struct {
	Repo *repository.AdminRepository
}

func NewAdminService(repo *repository.AdminRepository) *AdminService {
	return &AdminService{Repo: repo}
}

func (s *AdminService) Register(admin domain.Admin) error {
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
