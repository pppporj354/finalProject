package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type SocialMedia struct {
	gorm.Model
	Id             int    `json:"id"`
	Name           string `gorm:"not null" json:"name" validate:"required" json:"name"`
	SocialMediaUrl string `gorm:"not null" json:"social_media_url" validate:"required" json:"social_media_url"`
	UserId         int    `json:"user_id"`
	User           *User  `gorm:"foreignKey:UserId"`
}

func init() {
	validate = validator.New()
}

func (s *SocialMedia) Validate() error {
	return validate.Struct(s)
}
