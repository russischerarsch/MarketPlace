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
func (u *UserRepository) Create(ctx context.Context, user *users.User) error {
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
func (u UserRepository) GetAll(ctx context.Context) (*[]users.User, error) {
	SQLquery := `
	SELECT id, name, email, created_at FROM users
	`
	rows, err := u.db.Query(ctx, SQLquery)
	if err != nil {
		return &[]users.User{}, err
	}
	defer rows.Close()
	usersList := []users.User{}
	for rows.Next() {
		var user users.User
		err := rows.Scan(&user.ID, &user.FullName, &user.Email, &user.CreatedAt)
		if err != nil {
			return &[]users.User{}, err
		}
		usersList = append(usersList, user)
	}
	return &usersList, nil
}
func (u UserRepository) GetByID(ctx context.Context, id int) (*users.User, error) {
	SQLquery := `
	SELECT id, name, email, created_at FROM users
	WHERE id = $1
	`
	var user users.User
	if err := u.db.QueryRow(ctx, SQLquery, id).Scan(&user.ID, &user.FullName, &user.Email, &user.CreatedAt); err != nil {
		return &users.User{}, err
	}
	return &user, nil
}
func (u UserRepository) GetByEmail(ctx context.Context, email string) (*users.User, error) {
	SQLquery := `
	SELECT id, name, password, email, created_at FROM users
	WHERE email = $1
	`
	var user users.User
	err := u.db.QueryRow(ctx, SQLquery, email).Scan(&user.ID, &user.FullName, &user.PasswordHash, &user.Email, &user.CreatedAt)
	return &user, err
}
