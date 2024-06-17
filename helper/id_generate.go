package helper

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
)

func GenerateID(db *gorm.DB, tableName string, prefix string) (string, error) {
	var count int64
	if err := db.Table(tableName).Count(&count).Error; err != nil {
		return "", err
	}
	newID := count + 1
	return fmt.Sprintf("%s-%03d", strings.ToUpper(prefix), newID), nil
}
