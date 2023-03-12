package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"e-shop-backend/src/handlers"
	"e-shop-backend/src/repositories"
	"e-shop-backend/src/services"
)

func RegisteCategoryRoutes(router *gin.Engine, db *gorm.DB) {

	categoryRepo := repositories.NewCategoryRepo(db)
	categoryService := services.CategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	categoryRoutes := router.Group("category")
	{
		categoryRoutes.GET("", categoryHandler.GetAllCategories)
		categoryRoutes.POST("", categoryHandler.CreateCategory)
		categoryRoutes.PUT("/:id", categoryHandler.UpdateCategory)
		categoryRoutes.DELETE("/:id", categoryHandler.DeleteCategory)

	}

}
