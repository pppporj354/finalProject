package repositories

import "gorm.io/gorm"

type UserRepository struct {
	DB *gorm.DB
}
