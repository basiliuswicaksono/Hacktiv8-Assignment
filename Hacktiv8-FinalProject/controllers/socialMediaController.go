package controllers

import (
	"finalProject/models"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SocialMediaController struct {
	db *gorm.DB
}

func NewSocialMediaController(db *gorm.DB) *SocialMediaController {
	return &SocialMediaController{db: db}
}

func (s *SocialMediaController) CreateSocialMedia(ctx *gin.Context) {
	var (
		socialMedia models.SocialMedia
		err         error
	)

	userIDToken, _ := ctx.Get("id")
	userID := uint(userIDToken.(float64))

	err = ctx.ShouldBindJSON(&socialMedia)
	if err != nil {
		badRequestJsonResponse(ctx, err.Error())
		return
	}

	_, err = govalidator.ValidateStruct(&socialMedia)
	if err != nil {
		badRequestJsonResponse(ctx, err.Error())
		return
	}

	socialMedia.UserID = userID
	err = s.db.Create(&socialMedia).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusCreated, gin.H{
		"id":               socialMedia.ID,
		"name":             socialMedia.Name,
		"social_media_url": socialMedia.Social_media_url,
		"user_id":          socialMedia.UserID,
		"created_at":       socialMedia.CreatedAt,
	})
}

func (s *SocialMediaController) GetAllSocialMedia(ctx *gin.Context) {
	var socialMedias []models.SocialMedia

	err := s.db.Preload("User", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("email", "username", "id")
	}).Find(&socialMedias).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusOK, gin.H{
		"social_medias": socialMedias,
	})
}

func (s *SocialMediaController) UpdateSocialMedia(ctx *gin.Context) {
	socialMediaID := ctx.Param("socialMediaId")

	var (
		socialMedia   models.SocialMedia
		socialMediaDB models.SocialMedia
		err           error
	)

	err = ctx.ShouldBindJSON(&socialMedia)
	if err != nil {
		badRequestJsonResponse(ctx, err.Error())
		return
	}

	_, err = govalidator.ValidateStruct(&socialMedia)
	if err != nil {
		badRequestJsonResponse(ctx, err.Error())
		return
	}

	err = s.db.First(&socialMediaDB, "id=?", socialMediaID).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			noDataJsonResponse(ctx, err.Error())
			return
		}
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	err = s.db.Model(&socialMediaDB).Updates(models.SocialMedia{Name: socialMedia.Name, Social_media_url: socialMedia.Social_media_url}).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusOK, gin.H{
		"id":               socialMediaDB.ID,
		"name":             socialMediaDB.Name,
		"social_media_url": socialMediaDB.Social_media_url,
		"user_id":          socialMediaDB.UserID,
		"updated_at":       socialMediaDB.UpdatedAt,
	})
}

func (s *SocialMediaController) DeleteSocialMedia(ctx *gin.Context) {
	socialMediaID := ctx.Param("socialMediaId")

	var (
		socialMediaDB models.SocialMedia
		err           error
	)

	err = s.db.First(&socialMediaDB, "id=?", socialMediaID).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			noDataJsonResponse(ctx, err.Error())
			return
		}
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	err = s.db.Delete(&socialMediaDB).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusOK, gin.H{
		"message": "Your social media has been successfully deleted",
	})
}
