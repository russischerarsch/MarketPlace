package repositories

import (
	"context"
	orderitem "mini-ozon/intern/models/orderItem"
	"mini-ozon/intern/models/orders"

	"github.com/jackc/pgx/v5"
)

type OrderRepository struct {
	db *pgx.Conn
}

func CreateOrderRep(conn *pgx.Conn) *OrderRepository {
	return &OrderRepository{
		db: conn,
	}
}
func (u OrderRepository) Create(ctx context.Context, order *orders.Order) error {
	SQLquery := `
	INSERT INTO orders (user_id, status, created_at,total)
	VALUES($1,$2,$3,$4)
	RETURNING id
	`
	err := u.db.QueryRow(ctx, SQLquery, order.UserID, order.Status, order.CreatedAt, order.Total).Scan(&order.ID)
	if err != nil {
		return err
	}
	return nil
}
func (u OrderRepository) GetAll(ctx context.Context) ([]orders.Order, error) {
	SQLquery := `
	SELECT id, user_id, status, created_at FROM orders
	`
	rows, err := u.db.Query(ctx, SQLquery)
	if err != nil {
		return []orders.Order{}, err
	}
	orderList := []orders.Order{}
	for rows.Next() {
		var order orders.Order
		rows.Scan(&order.ID, &order.UserID, &order.Status, &order.CreatedAt)
		orderList = append(orderList, order)
	}
	return orderList, nil
}
func (u OrderRepository) GetByID(ctx context.Context, id int) (orders.Order, error) {
	SQLquery := `
	SELECT id, user_id, status, created_at FROM orders
	WHERE id = $1
	`
	var order orders.Order
	err := u.db.QueryRow(ctx, SQLquery, id).Scan(&order.ID, &order.UserID, &order.Status, &order.CreatedAt)
	if err != nil {
		return orders.Order{}, err
	}
	return order, nil
}
func (u OrderRepository) GetAllByUserId(ctx context.Context, id int) ([]orders.Order, error) {
	// 1. Получаем все заказы пользователя
	ordersQuery := `
        SELECT id, status, total, created_at
        FROM orders
        WHERE user_id = $1
    `
	rows, err := u.db.Query(ctx, ordersQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []orders.Order
	for rows.Next() {
		var o orders.Order
		if err := rows.Scan(&o.ID, &o.Status, &o.Total, &o.CreatedAt); err != nil {
			return nil, err
		}
		result = append(result, o)
	}
	return result, nil
}

func (r *OrderRepository) GetOrderItems(ctx context.Context, orderID int) ([]orderitem.OrderItem, error) {
	query := `
        SELECT id, product_id, quantity
        FROM order_items
        WHERE order_id = $1
    `
	rows, err := r.db.Query(ctx, query, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []orderitem.OrderItem
	for rows.Next() {
		var item orderitem.OrderItem
		if err := rows.Scan(&item.ID, &item.Product_id, &item.Quantity); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}
