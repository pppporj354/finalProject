package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gram/models"
	"gram/services"
	"io"
	"log"
	"net/http"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(usersService *services.UserService) *UserController {
	return &UserController{userService: usersService}
}

func (u *UserController) GetUsers(ctx *gin.Context) {
	users, respErr := u.userService.GetUsers()
	if respErr.Message != "" {
		log.Println("Error while getting users", respErr)
		ctx.AbortWithStatusJSON(respErr.Status, gin.H{"error": respErr.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (u *UserController) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	response, err := u.userService.GetUser(id)
	if err != nil {
		log.Println("Error while getting user", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (u *UserController) CreateUser(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return

	}
	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("Error while unmarshalling request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return

	}
	response, err := u.userService.CreateUser(&user)
	if err != nil {
		log.Println("Error while creating user", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, response)
}

func (u *UserController) UpdateUser(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("Error while unmarshalling request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	response, err := u.userService.UpdateUser(&user)
	if err != nil {
		log.Println("Error while updating user", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (u *UserController) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	ok, respErr := u.userService.DeleteUser(id)
	if !ok {
		log.Println("Error while deleting user", respErr)
		ctx.AbortWithStatusJSON(respErr.Status, gin.H{"error": respErr.Message})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
