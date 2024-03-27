package models

import (
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&User{}, &Photo{}, &Comment{}, &SocialMedia{})
	if err != nil {
		return err
	}
	return nil
}
