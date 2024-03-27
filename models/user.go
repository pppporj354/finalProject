package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id          int         `json:"id"`
	Username    string      `gorm:"not null;unique" validate:"required" json:"username"`
	Password    string      `gorm:"not null" validate:"required,min=6" json:"password"`
	Email       string      `gorm:"not null;unique" validate:"required,email" json:"email"`
	Age         int         `validate:"required,min=8" json:"age"`
	Photos      []Photo     `gorm:"foreignKey:UserId"`
	Comments    []Comment   `gorm:"foreignKey:UserId"`
	SocialMedia SocialMedia `gorm:"foreignKey:UserId"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func (u *User) Validate() error {
	return validate.Struct(u)
}
