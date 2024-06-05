package repository

import "mikti-depublic/model/domain"

type UserRepository interface {
	FindUserByEmail(email string) (*domain.User, error)
	FindAdminByEmail(email string) (*domain.Admin, error)
}