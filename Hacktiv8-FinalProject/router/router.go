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

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", userController.Register)
		userRouter.POST("/login", userController.Login)
		// with auth
		userRouter.Use(middlewares.Auth())
		userRouter.PUT("/", userController.UpdateUser)
		userRouter.DELETE("/", userController.DeleteUser)
	}

	return router
}
