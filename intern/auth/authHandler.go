package auth

import "github.com/gin-gonic/gin"

type AuthHandlers struct {
	AuthService *AuthService
}

func CreateAuthHandlers(service *AuthService) *AuthHandlers {
	return &AuthHandlers{
		AuthService: service,
	}
}

func (a AuthHandlers) Login(c *gin.Context) {
	var request struct {
		Email string `json:"email"`
		Pass  string `json:"password"`
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	token, err := a.AuthService.Login(request.Email, request.Pass, c)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"token": token})
}
