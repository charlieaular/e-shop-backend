package routes

import (
	"e-shop-backend/src/handlers"
	"e-shop-backend/src/middlewares"
	"e-shop-backend/src/repositories"
	"e-shop-backend/src/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewAddressRoutes(router *gin.Engine, db *gorm.DB) {
	addressRepo := repositories.NewAddressRepo(db)
	addressService := services.NewAddressService(addressRepo)
	addressHandler := handlers.NewAddressHandler(addressService)

	addressRoutes := router.Group("addresses").Use(middlewares.AuthMiddleware())
	{
		addressRoutes.POST("", addressHandler.CreateAddress)
		addressRoutes.GET("", addressHandler.GetAllAddresses)
		addressRoutes.PUT("/:id", addressHandler.UpdateAddress)
		addressRoutes.DELETE("/:id", addressHandler.DeleteAddress)
	}
}
