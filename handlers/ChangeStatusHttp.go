package handlers

import (
	"mini-ozon/intern/models/orders"
	"mini-ozon/intern/services"

	"github.com/gin-gonic/gin"
)

type ChangeStatusHandler struct {
	service *services.ChangeStatServ
}

func NewChangeHandler(service *services.ChangeStatServ) *ChangeStatusHandler {
	return &ChangeStatusHandler{
		service: service,
	}
}
func (ch ChangeStatusHandler) ChangeStatus(c *gin.Context) {

	var req struct {
		Order_id int           `json:"order_id"`
		Status   orders.Status `json:"status"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if err := ch.service.ChangeStatus(c, req.Order_id, req.Status); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "successfuly updated"})
}
