package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"e-shop-backend/src/services"
	"e-shop-backend/src/shared"
)

type productHandler struct {
	productService services.ProductService
}

func NewProductHandler(productService services.ProductService) productHandler {
	return productHandler{productService: productService}
}

func (h *productHandler) GetProductsByCategory(c *gin.Context) {

	categoryId, err := strconv.Atoi(c.Param("category"))

	if shared.HandleError(c, err) {
		return
	}

	products, err := h.productService.GetProductsByCategory(categoryId)

	if shared.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   true,
		"products": products,
	})
}

func (h *productHandler) GetProductById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if shared.HandleError(c, err) {
		return
	}

	product, err := h.productService.GetById(id)

	if shared.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"product": product,
	})
}
