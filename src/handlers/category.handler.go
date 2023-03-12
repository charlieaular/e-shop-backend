package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"e-shop-backend/src/models"
	createRequest "e-shop-backend/src/requests/category/create"
	updateRequest "e-shop-backend/src/requests/category/update"
	"e-shop-backend/src/services"
	"e-shop-backend/src/shared"
)

type categoryHandler struct {
	categoryService services.CategoryService
}

func NewCategoryHandler(categoryService services.CategoryService) categoryHandler {
	return categoryHandler{categoryService: categoryService}
}

func (h *categoryHandler) GetAllCategories(c *gin.Context) {

	categories, err := h.categoryService.GetAll()

	if shared.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":     true,
		"categories": categories,
	})
}

func (h *categoryHandler) CreateCategory(c *gin.Context) {

	var createCategoryRequest createRequest.CreateCategoryRequest

	err := c.ShouldBind(&createCategoryRequest)

	if shared.HandleError(c, err) {
		return
	}

	newCategory := models.Category{
		Name: createCategoryRequest.Name,
	}

	category, err := h.categoryService.CreateCategory(newCategory)

	if shared.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   true,
		"category": category,
	})

}

func (h *categoryHandler) UpdateCategory(c *gin.Context) {

	var updateCategoryRequest updateRequest.UpdateCategoryRequest

	err := c.ShouldBind(&updateCategoryRequest)

	if shared.HandleError(c, err) {
		return
	}

	categoryId, err := strconv.Atoi(c.Param("id"))

	if shared.HandleError(c, err) {
		return
	}

	newCategory := models.Category{
		Model: gorm.Model{
			ID: uint(categoryId),
		},
		Name: updateCategoryRequest.Name,
	}

	category, err := h.categoryService.UpdateCategory(newCategory)

	if shared.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   true,
		"category": category,
	})

}

func (h *categoryHandler) DeleteCategory(c *gin.Context) {

	categoryId := c.Param("id")

	err := h.categoryService.DeleteCategory(categoryId)

	if shared.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})

}
