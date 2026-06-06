package dbconnection

import (
	"context"

	"github.com/jackc/pgx/v5"
)

var conn *pgx.Conn
var Ctx context.Background()