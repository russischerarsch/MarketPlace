package handlers

import (
	orderitem "mini-ozon/intern/models/orderItem"
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
	var orderRequest orderitem.OrderRequest
	if err := c.BindJSON(&orderRequest); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	userId := c.GetInt("user_id")
	order, err := o.service.CreateOrder(c, userId, orderRequest.Items)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
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
func (o OrderHandlers) GetAllByUserId(c *gin.Context) {
	id := c.GetInt("id")
	items, err := o.service.GetByUserID(c, id)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
	}
	c.JSON(200, items)
}
