package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gram/models"
	"gram/repositories"
	"net/http"
)

type UserService struct {
	userRepository    *repositories.UserRepository
	socialRepository  *repositories.SocialRepository
	photoRepository   *repositories.PhotoRepository
	commentRepository *repositories.CommentRepository
}

func NewUserService(userRepository *repositories.UserRepository, socialRepository *repositories.SocialRepository, photoRepository *repositories.PhotoRepository, commentRepository *repositories.CommentRepository) *UserService {
	return &UserService{userRepository: userRepository, socialRepository: socialRepository, photoRepository: photoRepository, commentRepository: commentRepository}
}

func (u *UserService) CreateUser(user *models.User) (*models.User, models.ResponseError) {
	// Validate the user input
	if err := user.Validate(); err != nil {
		return nil, models.ResponseError{Message: err.Error(), Status: http.StatusBadRequest}
	}

	// Check if the user already exists
	var existingUser models.User
	result := u.userRepository.DB.Where("email = ?", user.Email).First(&existingUser)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, models.ResponseError{Message: "Failed to check if user exists", Status: http.StatusInternalServerError}
		}
	} else {
		return nil, models.ResponseError{Message: "User already exists", Status: http.StatusConflict}
	}

	// Hash the user's password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, models.ResponseError{Message: "Failed to hash password", Status: http.StatusInternalServerError}
	}
	user.Password = string(hashedPassword)

	// Save the new user to the database
	newUser, err := u.userRepository.Create(user)
	if err != nil {
		return nil, models.ResponseError{Message: "Failed to create user", Status: http.StatusInternalServerError}
	}

	return newUser, models.ResponseError{}
}

func (u *UserService) GetUsers() ([]models.User, models.ResponseError) {
	users, err := u.userRepository.GetAll()
	if err != nil {
		return nil, models.ResponseError{Message: "Failed to get users", Status: http.StatusInternalServerError}
	}
	return users, models.ResponseError{}
}

func (u *UserService) GetUser(id string) (*models.User, models.ResponseError) {
	user, err := u.userRepository.GetByID(id)
	if err != nil {
		return nil, models.ResponseError{Message: "Failed to get user", Status: http.StatusInternalServerError}
	}
	return user, models.ResponseError{}
}

func (u *UserService) UpdateUser(user *models.User) (*models.User, models.ResponseError) {
	// Validate the user input
	if err := user.Validate(); err != nil {
		return nil, models.ResponseError{Message: err.Error(), Status: http.StatusBadRequest}
	}

	// Check if the user exists
	var existingUser models.User
	result := u.userRepository.DB.Where("id = ?", user.ID).First(&existingUser)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, models.ResponseError{Message: "User does not exist", Status: http.StatusNotFound}
		}
		return nil, models.ResponseError{Message: "Failed to check if user exists", Status: http.StatusInternalServerError}
	}

	// Hash the user's password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, models.ResponseError{Message: "Failed to hash password", Status: http.StatusInternalServerError}
	}
	user.Password = string(hashedPassword)

	// Update the user in the database
	updatedUser, err := u.userRepository.Update(user)
	if err != nil {
		return nil, models.ResponseError{Message: "Failed to update user", Status: http.StatusInternalServerError}
	}

	return updatedUser, models.ResponseError{}
}

func (u *UserService) DeleteUser(id string) (bool, models.ResponseError) {
	// Check if the user exists
	var existingUser models.User
	result := u.userRepository.DB.Where("id = ?", id).First(&existingUser)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, models.ResponseError{Message: "User does not exist", Status: http.StatusNotFound}
		}
		return false, models.ResponseError{Message: "Failed to check if user exists", Status: http.StatusInternalServerError}
	}

	// Delete the user from the database
	err := u.userRepository.Delete(id)
	if err != nil {
		return false, models.ResponseError{Message: "Failed to delete user", Status: http.StatusInternalServerError}
	}

	return true, models.ResponseError{}
}

func (u *UserService) UserLogin(email, username, password string) (bool, models.ResponseError) {
	// Check if the user exists
	var existingUser models.User
	result := u.userRepository.DB.Where("email = ? and username = ?", email, username).First(&existingUser)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, models.ResponseError{Message: "User does not exist", Status: http.StatusNotFound}
		}
		return false, models.ResponseError{Message: "Failed to check if user exists", Status: http.StatusInternalServerError}
	}

	// Check if the password is correct
	err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(password))
	if err != nil {
		return false, models.ResponseError{Message: "Incorrect password", Status: http.StatusUnauthorized}
	}

	return true, models.ResponseError{}
}
