package requests

import "e-shop-backend/src/models"

type loginRespose struct {
	Status bool         `json:"status"`
	Token  string       `json:"token"`
	User   userResponse `json:"user"`
}

type userResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func LoginResponse(user models.User, token string) loginRespose {

	userResponse := userResponse{
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}

	loginResponse := loginRespose{
		Status: true,
		Token:  token,
		User:   userResponse,
	}

	return loginResponse

}
