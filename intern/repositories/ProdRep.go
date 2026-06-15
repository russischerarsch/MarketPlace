package repositories

import (
	"context"
	"mini-ozon/intern/models/products"

	"github.com/jackc/pgx/v5"
)

type ProductRepository struct {
	db *pgx.Conn
}

func CreateProdRep(conn *pgx.Conn) *ProductRepository {
	return &ProductRepository{
		db: conn,
	}
}
func (p ProductRepository) Create(ctx context.Context, prod *products.Product) error {
	SQLquery := `
	INSERT INTO products (name,price,descriptions,created_at)
	VALUES ($1,$2,$3,$4)
	RETURNING id
	`
	err := p.db.QueryRow(ctx, SQLquery, prod.Name, prod.Price, prod.Description, prod.Created_at).Scan(&prod.ID)
	if err != nil {
		return err
	}
	return nil
}
func (p ProductRepository) GetAll(ctx context.Context) ([]products.Product, error) {
	SQLquery := `
	SELECT id, name, price, descriptions, created_at FROM products
	ORDER BY id
	`
	rows, err := p.db.Query(ctx, SQLquery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var prods []products.Product
	for rows.Next() {
		var prod products.Product
		if err := rows.Scan(&prod.ID, &prod.Name, &prod.Price, &prod.Description, &prod.Created_at); err != nil {
			return nil, err
		}
		prods = append(prods, prod)
	}
	return prods, nil
}
func (p ProductRepository) GetByID(ctx context.Context, id int) (products.Product, error) {
	SQLquery := `
	SELECT id, name, descriptions, created_at, price FROM products
	WHERE id = $1
	`
	var product products.Product
	err := p.db.QueryRow(ctx, SQLquery, id).Scan(&product.ID, &product.Name, &product.Description, &product.Created_at, &product.Price)
	if err != nil {
		return products.Product{}, err
	}
	return product, nil
}
func (p ProductRepository) DeleteByID(ctx context.Context, id int) error {
	SQLquery := `
	DELETE FROM products 
	WHERE id = $1
	`
	_, err := p.db.Exec(ctx, SQLquery, id)

	return err

}
