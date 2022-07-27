package router

import (
	"finalProject/controllers"
	"finalProject/database"
	"finalProject/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	db := database.ConnectDB()
	userController := controllers.NewUserController(db)
	photoController := controllers.NewPhotoController(db)
	commentController := controllers.NewCommentController(db)
	socialMediaController := controllers.NewSocialMediaController(db)

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", userController.Register)
		userRouter.POST("/login", userController.Login)
		// with auth
		userRouter.Use(middlewares.Auth())
		userRouter.PUT("/", userController.UpdateUser)
		userRouter.DELETE("/", userController.DeleteUser)
	}

	photoRouter := router.Group("/photos")
	{
		photoRouter.Use(middlewares.Auth())
		photoRouter.POST("/", photoController.CreatePhoto)
		photoRouter.GET("/", photoController.GetAllPhotos)
		// with authorization
		photoRouter.Use(middlewares.PhotoAuthorization())
		photoRouter.PUT("/:photoId", photoController.UpdatePhoto)
		photoRouter.DELETE("/:photoId", photoController.DeletePhoto)
	}

	commentRouter := router.Group("/comments")
	{
		commentRouter.Use(middlewares.Auth())
		commentRouter.POST("/", commentController.CreateComment)
		commentRouter.GET("/", commentController.GetAllComments)
		//with authorization
		commentRouter.Use(middlewares.CommentAuthorization())
		commentRouter.PUT("/:commentId", commentController.UpdateComment)
		commentRouter.DELETE("/:commentId", commentController.DeleteComment)
	}

	socialMediaRouter := router.Group("/socialmedias")
	{
		socialMediaRouter.Use(middlewares.Auth())
		socialMediaRouter.POST("/", socialMediaController.CreateSocialMedia)
		socialMediaRouter.GET("/", socialMediaController.GetAllSocialMedia)
		// //with authorization
		socialMediaRouter.Use(middlewares.SocialMediaAuthorization())
		socialMediaRouter.PUT("/:socialMediaId", socialMediaController.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:socialMediaId", socialMediaController.DeleteSocialMedia)
	}

	return router
}
