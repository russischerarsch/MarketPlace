package repositories

import (
	"context"
	"mini-ozon/intern/models/users"

	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	db *pgx.Conn
}

func CreateUserRep(conn *pgx.Conn) *UserRepository {
	return &UserRepository{
		db: conn,
	}
}
func (u UserRepository) Create(ctx context.Context, user users.User) error {
	SQLquery := `
	INSERT INTO users (name, email, password, created_at)
	VALUES($1,$2,$3,$4)
	RETURNING id
	`
	err := u.db.QueryRow(ctx, SQLquery, user.FullName, user.Email, user.PasswordHash, user.CreatedAt).Scan(&user.ID)
	if err != nil {
		return err
	}
	return nil
}
func (u UserRepository) GetAll(ctx context.Context) ([]users.User, error) {
	SQLquery := `
	SELECT name, email, created_at FROM users
	`
	rows, err := u.db.Query(ctx, SQLquery)
	if err != nil {
		return err
	}
	defer rows.Close()
	users := []users.User{}
	for rows.Next() {
		var user users.User
		err := rows.Scan(&user.fullname)
		users = append(users)
	}
	return users, nil
}
