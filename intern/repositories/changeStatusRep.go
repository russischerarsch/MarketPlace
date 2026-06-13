package repositories

import (
	"context"
	"mini-ozon/intern/models/orders"

	"github.com/jackc/pgx/v5"
)

type ChangeRep struct {
	db *pgx.Conn
}

func CreateChangeRep(conn *pgx.Conn) *ChangeRep {
	return &ChangeRep{
		db: conn,
	}
}
func (c ChangeRep) ChangeStatus(status orders.Status, id int, ctx context.Context) error {
	SQLquery := `
	UPDATE orders SET status $1
	WHERE id = $2
	`
	if _, err := c.db.Exec(ctx, SQLquery, status, id); err != nil {
		return err
	}
	return nil
}
func (c ChangeRep) GetStatus(id int, ctx context.Context) (orders.Status, error) {
	SQLquery := `
	SELECT status FROM orders
	WHERE id = $2
	`
	var status orders.Status
	if err := c.db.QueryRow(ctx, SQLquery, status, id).Scan(&status); err != nil {
		return "", err
	}
	return status, nil
}
