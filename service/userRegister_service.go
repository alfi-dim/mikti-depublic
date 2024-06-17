package service

import (
	"mikti-depublic/helper"
	"mikti-depublic/model/domain"
	"mikti-depublic/repository"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRegisterService struct {
	Repo *repository.UserRegisterRepository
	DB   *gorm.DB
}

func NewUserRegisterService(repo *repository.UserRegisterRepository, db *gorm.DB) *UserRegisterService {
	return &UserRegisterService{Repo: repo, DB: db}
}

func (s *UserRegisterService) Register(user domain.User) error {
	userID, err := helper.GenerateID(s.DB, user.TableName(), "USER")
	if err != nil {
		return err
	}
	user.UserID = userID

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
