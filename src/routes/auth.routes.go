package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"e-shop-backend/src/handlers"
	"e-shop-backend/src/middlewares"
	"e-shop-backend/src/repositories"
	"e-shop-backend/src/services"
)

func RegisteAuthRoutes(router *gin.Engine, db *gorm.DB) {

	userRepo := repositories.NewUserRepo(db)
	userService := services.UserService(userRepo)
	authHandler := handlers.NewAuthHandler(userService)

	authRoutes := router.Group("auth")
	{
		authRoutes.POST("/login", authHandler.Login)
		authRoutes.POST("/register", authHandler.RegisterClient)
		authRoutes.GET("/user", middlewares.AuthMiddleware(), authHandler.GetCurrentUser)
	}

}
