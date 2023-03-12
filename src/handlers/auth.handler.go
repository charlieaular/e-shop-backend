package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"e-shop-backend/src/models"
	currentUserRequests "e-shop-backend/src/requests/auth/currentUser"
	loginRequests "e-shop-backend/src/requests/auth/login"
	registerRequests "e-shop-backend/src/requests/auth/register"
	"e-shop-backend/src/services"
	"e-shop-backend/src/shared"
)

type AuthHandler struct {
	UserService services.UserService
}

func NewAuthHandler(UserService services.UserService) AuthHandler {
	return AuthHandler{UserService: UserService}
}

func (authHandler *AuthHandler) Login(c *gin.Context) {
	var loginRequest loginRequests.LoginRequest

	err := c.ShouldBindJSON(&loginRequest)

	if shared.HandleError(c, err) {
		return
	}

	email := loginRequest.Email
	password := loginRequest.Password

	user, err := authHandler.UserService.GetUserByEmail(email)

	if shared.HandleError(c, err) {
		return
	}

	isSamePassword := shared.IsSamePassword([]byte(user.Password), []byte(password))

	if !isSamePassword {
		err := errors.New("usuario y/o contraseña incorrecto")
		if shared.HandleError(c, err) {
			return
		}
	}

	token := shared.GenerateToken(int(user.ID))

	loginResponse := loginRequests.LoginResponse(user, token)

	c.JSON(http.StatusOK, loginResponse)

}

func (authHandler *AuthHandler) GetCurrentUser(c *gin.Context) {

	id := c.GetInt("user_id")

	user, err := authHandler.UserService.GetUserById(id)

	if shared.HandleError(c, err) {
		return
	}

	currentUserResponse := currentUserRequests.CurrentUserResponse(user)

	c.JSON(http.StatusOK, currentUserResponse)

}

func (authHandler *AuthHandler) RegisterClient(c *gin.Context) {

	var registerRequests registerRequests.RegisterRequest

	err := c.ShouldBindJSON(&registerRequests)

	if shared.HandleError(c, err) {
		return
	}

	password := shared.HashPassword([]byte(registerRequests.Password))

	userModel := models.User{
		Name:     registerRequests.Name,
		Email:    registerRequests.Email,
		Password: password,
		Role:     "CLIENT",
	}

	user, err := authHandler.UserService.CreateClient(userModel)

	if shared.HandleError(c, err) {
		return
	}

	token := shared.GenerateToken(int(user.ID))

	loginResponse := loginRequests.LoginResponse(user, token)

	c.JSON(http.StatusOK, loginResponse)
}
