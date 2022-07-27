package controllers

import (
	"finalProject/models"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentController struct {
	db *gorm.DB
}

func NewCommentController(db *gorm.DB) *CommentController {
	return &CommentController{db: db}
}

func (c *CommentController) CreateComment(ctx *gin.Context) {
	var (
		comment models.Comment
		err     error
	)

	userIDToken, _ := ctx.Get("id")
	userID := uint(userIDToken.(float64))

	err = ctx.ShouldBindJSON(&comment)
	if err != nil {
		badRequestJsonResponse(ctx, err.Error())
		return
	}

	_, err = govalidator.ValidateStruct(&comment)
	if err != nil {
		badRequestJsonResponse(ctx, err.Error())
		return
	}

	comment.UserID = userID
	err = c.db.Create(&comment).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusCreated, gin.H{
		"id":         comment.ID,
		"message":    comment.Message,
		"photo_id":   comment.PhotoID,
		"user_id":    comment.UserID,
		"created_at": comment.CreatedAt,
	})
}

func (c *CommentController) GetAllComments(ctx *gin.Context) {
	var comments []models.Comment

	err := c.db.Preload("User", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("email", "username", "id")
	}).Preload("Photo", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("id", "title", "caption", "photo_url", "user_id")
	}).Find(&comments).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusOK, comments)
}

func (c *CommentController) UpdateComment(ctx *gin.Context) {
	commentID := ctx.Param("commentId")

	var (
		comment   models.Comment
		commentDB models.Comment
		err       error
	)

	err = ctx.ShouldBindJSON(&comment)
	if err != nil {
		badRequestJsonResponse(ctx, err.Error())
		return
	}

	_, err = govalidator.ValidateStruct(&comment)
	if err != nil {
		badRequestJsonResponse(ctx, err.Error())
		return
	}

	err = c.db.First(&commentDB, "id=?", commentID).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			noDataJsonResponse(ctx, err.Error())
			return
		}
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	err = c.db.Model(&commentDB).Updates(models.Comment{Message: comment.Message}).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusOK, gin.H{
		"id":         commentDB.ID,
		"message":    commentDB.Message,
		"user_id":    commentDB.UserID,
		"photo_id":   commentDB.PhotoID,
		"updated_at": commentDB.UpdatedAt,
	})
}

func (c *CommentController) DeleteComment(ctx *gin.Context) {
	commentID := ctx.Param("commentId")

	var (
		commentDB models.Comment
		err       error
	)

	err = c.db.First(&commentDB, "id=?", commentID).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			noDataJsonResponse(ctx, err.Error())
			return
		}
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	err = c.db.Delete(&commentDB).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}
