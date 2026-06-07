package handlers

import (
	"mini-ozon/intern/models/services"

	"github.com/gin-gonic/gin"
)

type UserHandlers struct {
	service *services.UserService
}

func CreateUserHandlers(service *services.UserService) *UserHandlers {
	return &UserHandlers{
		service: service,
	}
}
func (u UserHandlers) Create(c *gin.Context) {
	var request struct {
		FullName     string `json:"fullname"`
		Email        string `json:"email"`
		PasswordHash string `json:"password"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(500, gin.H{err.Error()})
	}
}
