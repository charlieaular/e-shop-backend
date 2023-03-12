package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"e-shop-backend/src/models"
	createRequest "e-shop-backend/src/requests/tag/create"
	updateRequest "e-shop-backend/src/requests/tag/update"

	"e-shop-backend/src/services"
	"e-shop-backend/src/shared"
)

type tagHandler struct {
	tagService services.TagService
}

func NewTagHandler(tagService services.TagService) tagHandler {
	return tagHandler{tagService: tagService}
}

func (h tagHandler) GetAllTags(c *gin.Context) {

	tags, err := h.tagService.GetAllTags()

	if shared.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"tags":   tags,
	})

}

func (h *tagHandler) CreateTag(c *gin.Context) {

	var createTagRequest createRequest.CreateTagRequest

	err := c.ShouldBind(&createTagRequest)

	if shared.HandleError(c, err) {
		return
	}

	newTag := models.Tag{
		Name: createTagRequest.Name,
	}

	tag, err := h.tagService.CreateTag(newTag)

	if shared.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"tag":    tag,
	})

}

func (h *tagHandler) UpdateTag(c *gin.Context) {

	var updateTagRequest updateRequest.UpdateTagRequest

	err := c.ShouldBind(&updateTagRequest)

	if shared.HandleError(c, err) {
		return
	}

	tagId, err := strconv.Atoi(c.Param("id"))

	if shared.HandleError(c, err) {
		return
	}

	newTag := models.Tag{
		Model: gorm.Model{
			ID: uint(tagId),
		},
		Name: updateTagRequest.Name,
	}

	tag, err := h.tagService.UpdateTag(newTag)

	if shared.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"tag":    tag,
	})

}

func (h *tagHandler) DeleteTag(c *gin.Context) {

	tagId := c.Param("id")

	err := h.tagService.DeleteTag(tagId)

	if shared.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})

}
