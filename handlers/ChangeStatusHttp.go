package handlers

import (
	"context"
	"encoding/json"
	"mini-ozon/intern/models/orders"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StatusHandler interface {
	ChangeStatus(ctx context.Context, id int, status orders.Status) error
}

type ChangeStatusHandler struct {
	service StatusHandler
}

func NewChangeHandler(service StatusHandler) ChangeStatusHandler {
	return ChangeStatusHandler{
		service: service,
	}
}

type Req struct {
	OrderID int           `json:"order_id"`
	Status  orders.Status `json:"status"`
}

func (ch ChangeStatusHandler) ChangeStatusGIN(c *gin.Context) {
	if c.Request.Method != http.MethodPatch {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "method patch required"})
		return
	}
	var req Req
	if err := c.BindJSON(&req); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if err := ch.service.ChangeStatus(c, req.OrderID, req.Status); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "successfuly updated"})
}

type Respone struct {
	Message string `json:"message"`
}

func (ch ChangeStatusHandler) ChangeStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "method patch required", http.StatusMethodNotAllowed)
		return
	}
	var req Req
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "failed to parse data from JSON", http.StatusBadRequest)
		return
	}
	if err := ch.service.ChangeStatus(r.Context(), req.OrderID, req.Status); err != nil {
		http.Error(w, "failed to update status, "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Respone{Message: "successfull update"})
}
