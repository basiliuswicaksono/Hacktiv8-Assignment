package controllers

import (
	"finalProject/models"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PhotoController struct {
	db *gorm.DB
}

func NewPhotoController(db *gorm.DB) *PhotoController {
	return &PhotoController{db: db}
}

func (p *PhotoController) CreatePhoto(ctx *gin.Context) {
	var (
		photo models.Photo
		err   error
	)

	userIDToken, _ := ctx.Get("id")
	userID := uint(userIDToken.(float64))

	err = ctx.ShouldBindJSON(&photo)
	if err != nil {
		badRequestJsonResponse(ctx, err.Error())
		return
	}

	_, err = govalidator.ValidateStruct(&photo)
	if err != nil {
		badRequestJsonResponse(ctx, err.Error())
		return
	}

	photo.UserID = userID
	err = p.db.Create(&photo).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusCreated, gin.H{
		"id":         photo.ID,
		"title":      photo.Title,
		"caption":    photo.Caption,
		"photo_url":  photo.Photo_url,
		"user_id":    photo.UserID,
		"created_at": photo.CreatedAt,
	})
}

func (p *PhotoController) GetAllPhotos(ctx *gin.Context) {
	var photos []models.Photo

	err := p.db.Preload("User", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("email", "username", "id")
	}).Find(&photos).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusOK, photos)
}

func (p *PhotoController) UpdatePhoto(ctx *gin.Context) {
	photoID := ctx.Param("photoId")

	var (
		photo   models.Photo
		photoDB models.Photo
		err     error
	)

	err = ctx.ShouldBindJSON(&photo)
	if err != nil {
		badRequestJsonResponse(ctx, err.Error())
		return
	}

	_, err = govalidator.ValidateStruct(&photo)
	if err != nil {
		badRequestJsonResponse(ctx, err.Error())
		return
	}

	err = p.db.First(&photoDB, "id=?", photoID).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			noDataJsonResponse(ctx, err.Error())
			return
		}
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	err = p.db.Model(&photoDB).Updates(photo).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusOK, gin.H{
		"id":         photoDB.ID,
		"title":      photoDB.Title,
		"caption":    photoDB.Caption,
		"photo_url":  photoDB.Photo_url,
		"user_id":    photoDB.UserID,
		"updated_at": photoDB.UpdatedAt,
	})
}

func (p *PhotoController) DeletePhoto(ctx *gin.Context) {
	photoID := ctx.Param("photoId")

	var (
		photoDB models.Photo
		err     error
	)

	err = p.db.First(&photoDB, "id=?", photoID).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			noDataJsonResponse(ctx, err.Error())
			return
		}
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	err = p.db.Delete(&photoDB).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}
