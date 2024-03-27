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

type PhotoController struct {
	photoService *services.PhotoService
}

func NewPhotoController(photoService *services.PhotoService) *PhotoController {
	return &PhotoController{photoService: photoService}
}

func (p *PhotoController) GetPhotos(ctx *gin.Context) {
	photos, err := p.photoService.GetPhotos()
	if err != nil {
		log.Println("Error while getting photos", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, photos)
}

func (p *PhotoController) GetPhoto(ctx *gin.Context) {
	id := ctx.Param("id")
	response, err := p.photoService.GetPhoto(id)
	if err != nil {
		log.Println("Error while getting photo", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (p *PhotoController) CreatePhoto(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return

	}
	var photo models.Photo
	err = json.Unmarshal(body, &photo)
	if err != nil {
		log.Println("Error while unmarshalling request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return

	}
	response, err := p.photoService.CreatePhoto(&photo)
	if err != nil {
		log.Println("Error while creating photo", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (p *PhotoController) UpdatePhoto(ctx *gin.Context) {
	id := ctx.Param("id")
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	var photo models.Photo
	err = json.Unmarshal(body, &photo)
	if err != nil {
		log.Println("Error while unmarshalling request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	response, err := p.photoService.UpdatePhoto(id, &photo)
	if err != nil {
		log.Println("Error while updating photo", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (p *PhotoController) DeletePhoto(ctx *gin.Context) {
	id := ctx.Param("id")
	err := p.photoService.DeletePhoto(id)
	if err != nil {
		log.Println("Error while deleting photo", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Photo deleted successfully"})
}
