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

type CommentController struct {
	commentService *services.CommentService
}

func NewCommentController(commentService *services.CommentService) *CommentController {
	return &CommentController{commentService: commentService}
}

func (c *CommentController) GetComments(ctx *gin.Context) {
	comments, err := c.commentService.GetComments()
	if err != nil {
		log.Println("Error while getting comments", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, comments)
}

func (c *CommentController) GetComment(ctx *gin.Context) {
	id := ctx.Param("id")
	response, err := c.commentService.GetComment(id)
	if err != nil {
		log.Println("Error while getting comment", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *CommentController) CreateComment(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return

	}
	var comment models.Comment
	err = json.Unmarshal(body, &comment)
	if err != nil {
		log.Println("Error while unmarshalling request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return

	}
	response, err := c.commentService.CreateComment(&comment)
	if err != nil {
		log.Println("Error while creating comment", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *CommentController) UpdateComment(ctx *gin.Context) {
	id := ctx.Param("id")
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return

	}
	var comment models.Comment
	err = json.Unmarshal(body, &comment)
	if err != nil {
		log.Println("Error while unmarshalling request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return

	}
	response, err := c.commentService.UpdateComment(id, &comment)
	if err != nil {
		log.Println("Error while updating comment", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *CommentController) DeleteComment(ctx *gin.Context) {
	id := ctx.Param("id")
	err := c.commentService.DeleteComment(id)
	if err != nil {
		log.Println("Error while deleting comment", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.Status(http.StatusNoContent)
}
