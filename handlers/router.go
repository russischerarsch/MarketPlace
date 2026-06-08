package handlers

import "github.com/gin-gonic/gin"

func SetupRouter(productHandler *ProductHandler, userUserHandler *UserHandlers, orderHandler *OrderHandlers) *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	v1 := r.Group("/api/v1")
	v1.POST("/products", productHandler.CreateHandler)
	v1.GET("/products", productHandler.GetAllHandler)
	v1.GET("/products/:id", productHandler.GetByIDHandler)
	v1.DELETE("/products/:id", productHandler.DeleteByIDhandler)

	v1.POST("/users", userUserHandler.CreateHandler)
	v1.GET("/users", userUserHandler.GetAllHandler)
	v1.GET("/users/:id", userUserHandler.GetByIDHandler)

	v1.POST("/orders", orderHandler.CreateOrderHandler)
	v1.GET("/orders", orderHandler.GetAllOrdersHandler)
	v1.GET("/orders/:id", orderHandler.GetByIdHandler)

	return r
}
