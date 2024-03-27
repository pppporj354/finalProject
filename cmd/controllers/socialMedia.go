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

type SocialController struct {
	socialService *services.SocialService
}

func NewSocialController(socialService *services.SocialService) *SocialController {
	return &SocialController{socialService: socialService}
}

func (s *SocialController) GetSocials(ctx *gin.Context) {
	socials, err := s.socialService.GetSocials()
	if err != nil {
		log.Println("Error while getting socials", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, socials)
}

func (s *SocialController) GetSocial(ctx *gin.Context) {
	id := ctx.Param("id")
	response, err := s.socialService.GetSocial(id)
	if err != nil {
		log.Println("Error while getting social ", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (s *SocialController) CreateSocial(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return

	}
	var social models.SocialMedia
	err = json.Unmarshal(body, &social)
	if err != nil {
		log.Println("Error while unmarshalling request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return

	}
	response, err := s.socialService.CreateSocial(&social)
	if err != nil {
		log.Println("Error while creating social", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (s *SocialController) UpdateSocial(ctx *gin.Context) {
	id := ctx.Param("id")
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	var social models.SocialMedia
	err = json.Unmarshal(body, &social)
	if err != nil {
		log.Println("Error while unmarshalling request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	response, err := s.socialService.UpdateSocial(id, &social)
	if err != nil {
		log.Println("Error while updating social", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (s *SocialController) DeleteSocial(ctx *gin.Context) {
	id := ctx.Param("id")
	err := s.socialService.DeleteSocial(id)
	if err != nil {
		log.Println("Error while deleting social", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Social deleted successfully"})
}
