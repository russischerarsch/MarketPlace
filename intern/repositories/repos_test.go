package repositories

import (
	"context"
	"mini-ozon/intern/models/products"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
)

func setupTestDB(t *testing.T) *pgx.Conn {
	connStr := "postgres://postgres:pass@localhost:5432/testdb?sslmode=disable"
	pool, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		t.Fatal(err)
	}
	return pool
}

func TestProductRepository_Create(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close(t.Context())

	repo := CreateProdRep(db)

	product := &products.Product{Name: "Test", Price: 100}
	err := repo.Create(context.Background(), product)

	assert.NoError(t, err)
	assert.NotZero(t, product.ID)
}
