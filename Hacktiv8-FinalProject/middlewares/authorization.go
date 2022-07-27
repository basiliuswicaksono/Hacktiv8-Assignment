package middlewares

import (
	"finalProject/database"
	"finalProject/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PhotoAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDB()
		photoID, err := strconv.Atoi(ctx.Param("photoId"))

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   err.Error(),
			})
			return
		}

		userIDToken, _ := ctx.Get("id")
		userID := uint(userIDToken.(float64))
		Photo := models.Photo{}

		err = db.Select("user_id").First(&Photo, photoID).Error
		if err != nil {
			if err.Error() == gorm.ErrRecordNotFound.Error() {
				ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"success": false,
					"error":   err.Error(),
				})
				return
			}

			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   err.Error(),
			})
			return
		}

		if Photo.UserID != userID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "UNAUTHORIZED - you are not allowed to access this data",
			})
			return
		}

		ctx.Next()
	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDB()
		commentID, err := strconv.Atoi(ctx.Param("commentId"))

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   err.Error(),
			})
			return
		}

		userIDToken, _ := ctx.Get("id")
		userID := uint(userIDToken.(float64))
		Comment := models.Comment{}

		err = db.Select("user_id").First(&Comment, commentID).Error
		if err != nil {
			if err.Error() == gorm.ErrRecordNotFound.Error() {
				ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"success": false,
					"error":   err.Error(),
				})
				return
			}

			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   err.Error(),
			})
			return
		}

		if Comment.UserID != userID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "UNAUTHORIZED - you are not allowed to access this data",
			})
			return
		}

		ctx.Next()
	}
}

func SocialMediaAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDB()
		socialMediaID, err := strconv.Atoi(ctx.Param("socialMediaId"))

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   err.Error(),
			})
			return
		}

		userIDToken, _ := ctx.Get("id")
		userID := uint(userIDToken.(float64))
		SocialMedia := models.SocialMedia{}

		err = db.Select("user_id").First(&SocialMedia, socialMediaID).Error
		if err != nil {
			if err.Error() == gorm.ErrRecordNotFound.Error() {
				ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"success": false,
					"error":   err.Error(),
				})
				return
			}

			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   err.Error(),
			})
			return
		}

		if SocialMedia.UserID != userID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "UNAUTHORIZED - you are not allowed to access this data",
			})
			return
		}

		ctx.Next()
	}
}
