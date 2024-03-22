package converter

import (
	"tot_golang/internal/entity"
	"tot_golang/internal/models"
)

func UserToResponse(user *entity.User) *models.UserResponse {
	return &models.UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Division: user.Division,
	}
}

func UsersToResponse(user []*entity.User) []models.UserResponse {
	var response []models.UserResponse

	for _, us := range user {
		response = append(response, *UserToResponse(us))
	}

	return response
}
