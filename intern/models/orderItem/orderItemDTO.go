package orderitem

import "mini-ozon/intern/models/orders"

type OrderItemRequest struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type OrderRequest struct {
	Items []OrderItemRequest `json:"items"`
}
type OrderWithItem struct {
	OrderItem []OrderItem
	Order     orders.Order
}
