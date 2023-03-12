package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"e-shop-backend/src/handlers"
	"e-shop-backend/src/repositories"
	"e-shop-backend/src/services"
)

func RegisterTagRoutes(router *gin.Engine, db *gorm.DB) {
	tagRepo := repositories.NewTagRepository(db)
	tagService := services.NewTagService(tagRepo)
	tagHandler := handlers.NewTagHandler(tagService)

	tagRoutes := router.Group("tags")
	{
		tagRoutes.GET("", tagHandler.GetAllTags)
		tagRoutes.POST("", tagHandler.CreateTag)
		tagRoutes.PUT("/:id", tagHandler.UpdateTag)
		tagRoutes.DELETE("/:id", tagHandler.DeleteTag)
	}

}
