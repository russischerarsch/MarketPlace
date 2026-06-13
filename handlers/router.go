package handlers

import (
	"mini-ozon/handlers"
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
	v1.GET("/orders", orderHandler.GetAllOrdersHandler)
	v1.GET("/users", userUserHandler.GetAllHandler)
	v1.GET("/users/:id", userUserHandler.GetByIDHandler)

	authorized := v1.Group("/").Use(auth.AuthMiddleWare())

	authorized.POST("/products", productHandler.CreateHandler)
	authorized.GET("/products", productHandler.GetAllHandler)
	authorized.GET("/products/:id", productHandler.GetByIDHandler)
	authorized.DELETE("/products/:id", productHandler.DeleteByIDhandler)

	authorized.POST("/orders", orderHandler.CreateOrderHandler)
	authorized.GET("/orders/:id", orderHandler.GetByIdHandler)
	authorized.PATCH("/orders", handlers.ChangeStatusHandler)
	authorized.GET("/orders/my", orderHandler.GetAllByUserId)

	authorized.POST("/orders/:id/items")
	authorized.GET("/orders/:id/items")

	return r
}
