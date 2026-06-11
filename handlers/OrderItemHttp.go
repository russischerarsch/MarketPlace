package handlers

import (
	"mini-ozon/intern/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderItemHandlers struct {
	service *services.OrderItemService
}

func CreateOrderItemHandler(service *services.OrderItemService) *OrderItemHandlers {
	return &OrderItemHandlers{
		service: service,
	}
}
func (o OrderItemHandlers) CreateOrderItemHandler(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	var request struct {
		Product_id int `json:"product_id"`
		Quantity   int `json:"quantity"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	order, err := o.service.CreateOrderItem(c, orderID, request.Product_id, request.Quantity)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, order)
}
func (o OrderItemHandlers) GetAllOrderItemHandlers(c *gin.Context) {
	orders, err := o.service.GetAll(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, orders)
}
func (o OrderItemHandlers) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	order, err := o.service.GetById(c, id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, order)
}
jwt.