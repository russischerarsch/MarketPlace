package handlers

import (
	"mini-ozon/intern/services"
	"strconv"

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
func (u UserHandlers) CreateHandler(c *gin.Context) {
	var request struct {
		FullName     string `json:"fullname"`
		Email        string `json:"email"`
		PasswordHash string `json:"password"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	user, err := u.service.CreateUser(c, request.FullName, request.Email, request.PasswordHash)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}
func (u UserHandlers) GetAllHandler(c *gin.Context) {
	usersList, err := u.service.GetAllService(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, usersList)
}
func (u UserHandlers) GetByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	user, err := u.service.GetByID(c, id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}
