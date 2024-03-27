package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Id      int    `json:"id"`
	UserId  int    `json:"user_id"`
	PhotoId int    `json:"photo_id"`
	Message string `gorm:"not null" validate:"required" json:"message"`
	User    User   `gorm:"foreignKey:UserId"`
	Photo   Photo  `gorm:"foreignKey:PhotoId"`
}

func init() {
	validate = validator.New()
}

func (c *Comment) Validate() error {
	return validate.Struct(c)
}
