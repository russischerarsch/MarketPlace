package repositories

import (
	"context"
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
	INSERT INTO orders (user_id, status, created_at)
	VALUES($1,$2,$3)
	RETURNING id
	`
	err := u.db.QueryRow(ctx, SQLquery, order.UserID, order.Status, order.CreatedAt).Scan(&order.ID)
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
