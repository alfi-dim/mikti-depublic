package service

import (
	"mikti-depublic/model/domain"
	"mikti-depublic/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserRegisterService struct {
	Repo *repository.UserRegisterRepository
}

func NewUserRegisterService(repo *repository.UserRegisterRepository) *UserRegisterService {
	return &UserRegisterService{Repo: repo}
}

func (s *UserRegisterService) Register(user domain.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	err = s.Repo.Save(user)
	if err != nil {
		return err
	}
	return nil
}
