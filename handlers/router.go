package handlers

import "github.com/gin-gonic/gin"

func SetupRouter(productHandler *ProductHandler) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	v1.POST("/products", productHandler.CreateHandler)
	v1.GET("/products", productHandler.GetAllHandler)
	v1.GET("/products/:id", productHandler.GetByIDHandler)
	v1.DELETE("/products/:id", productHandler.DeleteByIDhandler)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	return r
}
