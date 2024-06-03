package helper

import "gorm.io/gorm"

type GormOps interface {
	Create(data interface{}) error
}
type GormOpsImpl struct {
	db *gorm.DB
}

func NewGormOps(db *gorm.DB) *GormOpsImpl {
	return &GormOpsImpl{db: db}
}

func (ops *GormOpsImpl) Create(data interface{}) error {
	err := ops.db.Create(data).Error

	if err != nil {
		return err
	}

	return nil
}
