package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	Id       int       `json:"id"`
	Title    string    `gorm:"not null" json:"title" validate:"required" json:"title"`
	Caption  string    `gorm:"not null" json:"caption"`
	PhotoUrl string    `gorm:"not null" json:"photo_url" validate:"required" json:"photo_url"`
	UserId   int       `json:"user_id"`
	User     User      `gorm:"foreignKey:UserId"`
	Comments []Comment `gorm:"foreignKey:PhotoId"`
}

func init() {
	validate = validator.New()
}

func (p *Photo) Validate() error {
	return validate.Struct(p)
}
