package dbconnection

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func CheckConnection(ctx context.Context) (*pgx.Conn, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	connStr := os.Getenv("CONN_STRING")
	connection, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return nil, err
	}
	return connection, nil
}
