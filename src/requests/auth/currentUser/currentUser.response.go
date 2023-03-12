package requests

import "e-shop-backend/src/models"

type currentUserResponse struct {
	Status bool   `json:"status"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Role   string `json:"role"`
}

func CurrentUserResponse(user models.User) currentUserResponse {

	currentUserResponse := currentUserResponse{
		Status: true,
		Name:   user.Name,
		Email:  user.Email,
		Role:   user.Role,
	}

	return currentUserResponse

}
