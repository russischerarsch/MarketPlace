package dbconnection

import (
	"context"
	"mini-ozon/intern/products"

	"github.com/jackc/pgx/v5"
)

type ProductRepository struct {
	db *pgx.Conn
}

func (p *ProductRepository) Create(conn *pgx.Conn) *ProductRepository {
	return &ProductRepository{
		db: conn,
	}
}
func (p ProductRepository) Сreate(ctx context.Context, prod *products.Product) error {
	SQLquery := `
	INSERT INTO products (price,descriptions,created_at),
	VALUES ($1,$2,$3)
	RETURNING id
	`
	err := p.db.QueryRow(ctx, SQLquery, prod.Price, prod.Description, prod.Created_at).Scan(&prod.ID)
	if err != nil {
		return err
	}
	return nil
}
