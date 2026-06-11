package handlers

import (
	"mini-ozon/intern/auth"

	"github.com/gin-gonic/gin"
)

func SetupRouter(authHandler auth.AuthHandlers, productHandler *ProductHandler, userUserHandler *UserHandlers, orderHandler *OrderHandlers) *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	v1 := r.Group("/api/v1")
	v1.POST("/login", authHandler.Login)
	v1.POST("/register", userUserHandler.CreateHandler)

	authorized := v1.Group("/").Use(auth.AuthMiddleWare())

	authorized.POST("/products", productHandler.CreateHandler)
	authorized.GET("/products", productHandler.GetAllHandler)
	authorized.GET("/products/:id", productHandler.GetByIDHandler)
	authorized.DELETE("/products/:id", productHandler.DeleteByIDhandler)

	authorized.GET("/users", userUserHandler.GetAllHandler)
	authorized.GET("/users/:id", userUserHandler.GetByIDHandler)

	authorized.POST("/orders", orderHandler.CreateOrderHandler)
	authorized.GET("/orders", orderHandler.GetAllOrdersHandler)
	authorized.GET("/orders/:id", orderHandler.GetByIdHandler)

	authorized.POST("/orders/:id/items")
	authorized.GET("/orders/:id/items")

	return r
}
