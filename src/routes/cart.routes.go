package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"e-shop-backend/src/handlers"
	"e-shop-backend/src/repositories"
	"e-shop-backend/src/services"
)

func RegisteCartRoutes(router *gin.Engine, db *gorm.DB) {

	productRepo := repositories.NewProductRepo(db)
	productService := services.ProductService(productRepo)

	cartProductRepo := repositories.NewCartProductRepo(db)
	cartProductService := services.CartProductService(cartProductRepo)

	cartRepo := repositories.NewCartRepo(db)
	cartService := services.CartService(cartRepo)
	cartHandler := handlers.NewCartHandler(cartService, productService, cartProductService)

	cartRoutes := router.Group("cart")
	{
		cartRoutes.GET("/products", cartHandler.GetCartProducts)
		cartRoutes.DELETE("/products/:id", cartHandler.DeleteCartProduct)
		cartRoutes.POST("", cartHandler.AddProductToCart)
		cartRoutes.POST("/pay", cartHandler.PayCart)

	}

}
