package service

import (
	"errors"
	"mikti-depublic/helper"
	"mikti-depublic/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	repository   repository.UserRepository
	tokenUseCase helper.TokenUseCase
}

func NewUserService(repository repository.UserRepository, tokenUseCase helper.TokenUseCase) *UserServiceImpl {
	return &UserServiceImpl{
		repository:   repository,
		tokenUseCase: tokenUseCase,
	}
}

func (service *UserServiceImpl) LoginUser(email string, password string) (map[string]interface{}, error) {
	user, err := service.repository.FindUserByEmail(email)
	if err != nil {
		return nil, errors.New("email tidak ditemukan")
	}

	errPass := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errPass != nil {
		return nil, errors.New("password anda salah")
	}

	expiredTime := time.Now().Local().Add(1 * time.Hour)

	claims := helper.JwtCustomClaims{
		ID:    user.UserID,
		Name:  user.Name,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "kelompok-03",
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}

	token, errToken := service.tokenUseCase.GenerateAccessToken(claims)
	if errToken != nil {
		return nil, errors.New("ada kesalahan generate token")
	}

	return helper.ResponseToJson{"token": token}, nil

}

func (service *UserServiceImpl) LoginAdmin(email string, password string) (map[string]interface{}, error) {
	admin, err := service.repository.FindAdminByEmail(email)
	if err != nil {
		return nil, errors.New("email tidak ditemukan")
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password))
	if err != nil {
		return nil, errors.New("password anda salah")
	}

	expiredTime := time.Now().Local().Add(1 * time.Hour)
	claims := helper.JwtCustomClaims{
		ID:    admin.ID,
		Name:  admin.Name,
		Email: admin.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Kelompok-03",
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}

	token, errToken := service.tokenUseCase.GenerateAccessToken(claims)
	if errToken != nil {
		return nil, errors.New("ada kesalahan generate token")
	}

	return helper.ResponseToJson{"token": token}, nil

}
