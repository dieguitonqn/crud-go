package infrastructure

import (
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *User) error {
	result := db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
