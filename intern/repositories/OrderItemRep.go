package repositories

import (
	"context"
	orderitem "mini-ozon/intern/models/orderItem"

	"github.com/jackc/pgx/v5"
)

type OrderItemRepository struct {
	db *pgx.Conn
}

func CreateOrderItemRep(conn *pgx.Conn) *OrderItemRepository {
	return &OrderItemRepository{
		db: conn,
	}
}
func (o OrderItemRepository) BeginTx(ctx context.Context) (pgx.Tx, error) { return o.db.Begin(ctx) }
func (o *OrderItemRepository) CreateOrderItem(
	ctx context.Context,
	orderItem *orderitem.OrderItem,
	tx pgx.Tx) error {
	SQLquery := `
		INSERT INTO order_item (order_id, product_id, quantity)
		VALUES($1,$2,$3)
		RETURNING id
		`
	err := tx.QueryRow(ctx, SQLquery, orderItem.Order_id, orderItem.Product_id, orderItem.Quantity).Scan(&orderItem.ID)
	if err != nil {
		return err
	}
	return nil
}
func (o *OrderItemRepository) GetAllOrderItem(ctx context.Context) ([]orderitem.OrderItem, error) {
	SQLquery := `
	SELECT id, order_id, product_id, quantity FROM order_item
	`
	rows, err := o.db.Query(ctx, SQLquery)
	if err != nil {
		return []orderitem.OrderItem{}, err
	}
	orderItemList := []orderitem.OrderItem{}
	for rows.Next() {
		var order orderitem.OrderItem
		if err := rows.Scan(&order.ID, &order.Order_id, &order.Product_id, &order.Quantity); err != nil {
			return []orderitem.OrderItem{}, err
		}
		orderItemList = append(orderItemList, order)
	}
	return orderItemList, nil
}
func (o *OrderItemRepository) GetByIdOrderItem(ctx context.Context, id int) (*orderitem.OrderItem, error) {
	SQLquery := `
	SELECT id, order_id, product_id, quantity FROM order_item
	WHERE id = $1
	`
	var row orderitem.OrderItem
	err := o.db.QueryRow(ctx, SQLquery, id).Scan(&row.ID, &row.Order_id, &row.Product_id, &row.Quantity)
	if err != nil {
		return nil, err
	}
	return &row, nil
}
