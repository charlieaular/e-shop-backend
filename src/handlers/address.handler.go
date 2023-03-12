package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"e-shop-backend/src/models"
	createRequest "e-shop-backend/src/requests/address/create"
	updateRequest "e-shop-backend/src/requests/address/update"
	"e-shop-backend/src/services"
	"e-shop-backend/src/shared"
)

type AddressHandler struct {
	addressService services.AddressService
}

func NewAddressHandler(addressService services.AddressService) AddressHandler {
	return AddressHandler{addressService: addressService}
}

func (h AddressHandler) CreateAddress(c *gin.Context) {
	userId := c.GetInt("user_id")

	var createRequest createRequest.CreateAddressRequest

	err := c.ShouldBind(&createRequest)

	if shared.HandleError(c, err) {
		return
	}

	newAddress := models.Address{
		Address: createRequest.Address,
		UserID:  userId,
	}

	add, err := h.addressService.CreateAddress(newAddress)

	if shared.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"address": add,
	})

}

func (h AddressHandler) GetAllAddresses(c *gin.Context) {
	userId := c.GetInt("user_id")

	addresses, err := h.addressService.GetAllAddresses(userId)

	if shared.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":    true,
		"addresses": addresses,
	})

}

func (h AddressHandler) UpdateAddress(c *gin.Context) {

	var updateAddressRequest updateRequest.UpdateAddressRequest

	err := c.ShouldBind(&updateAddressRequest)

	if shared.HandleError(c, err) {
		return
	}

	addressId, err := strconv.Atoi(c.Param("id"))

	if shared.HandleError(c, err) {
		return
	}

	newAddress := models.Address{
		Model: gorm.Model{
			ID: uint(addressId),
		},
		Address: updateAddressRequest.Address,
	}

	address, err := h.addressService.UpdateAddress(newAddress)

	if shared.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"address": address,
	})

}

func (h AddressHandler) DeleteAddress(c *gin.Context) {

	addressId := c.Param("id")

	err := h.addressService.DeleteAddress(addressId)

	if shared.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})
}
