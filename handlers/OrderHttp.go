package handlers

import (
	"mini-ozon/intern/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderHandlers struct {
	service *services.OrderService
}

func CreateOrderHandlers(service *services.OrderService) *OrderHandlers {
	return &OrderHandlers{
		service: service,
	}
}
func (o OrderHandlers) CreateOrderHandler(c *gin.Context) {
	var request struct {
		User_id int `json:"user_id"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	order, err := o.service.CreateProduct(c, request.User_id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, order)
}
func (o OrderHandlers) GetAllOrdersHandler(c *gin.Context) {
	orders, err := o.service.GetAll(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, orders)
}

func (o OrderHandlers) GetByIdHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	order, err := o.service.GetByID(c, id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
	}
	c.JSON(200, order)
}
