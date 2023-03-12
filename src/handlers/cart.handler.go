package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"e-shop-backend/src/models"
	requests "e-shop-backend/src/requests/cart"
	"e-shop-backend/src/services"
	"e-shop-backend/src/shared"
)

type cartHandler struct {
	cartService        services.CartService
	productService     services.ProductService
	cartProductService services.CartProductService
}

func NewCartHandler(cartService services.CartService, productService services.ProductService, cartProductService services.CartProductService) cartHandler {
	return cartHandler{cartService: cartService, productService: productService, cartProductService: cartProductService}
}

func (h *cartHandler) GetCartProducts(c *gin.Context) {
	cartProducts, err := h.cartProductService.GetCartProducts()

	if shared.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":       true,
		"CartProducts": cartProducts,
	})
}

func (h *cartHandler) AddProductToCart(c *gin.Context) {
	var addProductToCartRequest requests.AddProductToCartRequest

	err := c.ShouldBind(&addProductToCartRequest)

	if shared.HandleError(c, err) {
		return
	}
	product, err := h.productService.GetById(addProductToCartRequest.ProductID)

	if shared.HandleError(c, err) {
		return
	}
	currentCart, err := h.cartService.CurrentActiveCart()

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		currentCart, err = h.cartService.CreateCart()
	}

	if shared.HandleError(c, err) {
		return
	}
	cartProduct := models.CartProduct{
		Name:        product.Name,
		Image:       product.Image,
		Description: product.Description,
		Price:       product.Price,
		Quantity:    1,
		CartID:      int(currentCart.ID),
	}

	newCartProduct, err := h.cartProductService.CreateCartProduct([]models.CartProduct{cartProduct})

	if shared.HandleError(c, err) {
		return
	}
	err1 := h.cartService.UpdateCartPrice()

	if shared.HandleError(c, err1) {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":         true,
		"newCartProduct": newCartProduct,
	})
}

func (h *cartHandler) DeleteCartProduct(c *gin.Context) {
	cartProductId, err := strconv.Atoi(c.Param("id"))

	if shared.HandleError(c, err) {
		return
	}
	status, err := h.cartProductService.DeleteCartProduct(cartProductId)

	if shared.HandleError(c, err) {
		return
	}
	cartProducts, err := h.cartProductService.GetCartProducts()

	if shared.HandleError(c, err) {
		return
	}
	if len(cartProducts) == 0 {
		err := h.cartService.DesactivateCart()

		if shared.HandleError(c, err) {
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status": status,
	})

}

func (h *cartHandler) PayCart(c *gin.Context) {

	err := h.cartService.PayCart()

	if shared.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})
}
