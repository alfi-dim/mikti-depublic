package service

import (
	"mikti-depublic/helper"
	"mikti-depublic/model/domain"
	"mikti-depublic/model/web"
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

func (s *UserRegisterService) Register(request web.UserRegisterRequest) (domain.User, error) {
	user := domain.User{}
	userID, err := helper.GenerateID(s.DB, user.TableName(), "USER")
	if err != nil {
		return domain.User{}, err
	}
	user.UserID = userID

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return domain.User{}, err
	}
	userReq := domain.User{
		UserID: userID,
		Name: request.Name,
		Username: request.Username,
		Email: request.Email,
		Password: string(hashedPassword),
		Phonenumber: request.PhoneNumber,
	}

	newUser, err := s.Repo.Save(userReq)
	if err != nil {
		return domain.User{}, err
	}
	return newUser, nil
}