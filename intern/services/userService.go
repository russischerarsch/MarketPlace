package services

import (
	"context"
	"errors"
	"mini-ozon/intern/models/users"
	"mini-ozon/intern/repositories"
	"strconv"
	"time"

	"github.com/patrickmn/go-cache"
)

type UserService struct {
	repo  *repositories.UserRepository
	cache *cache.Cache
}

func CreateUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{
		repo:  repo,
		cache: cache.New(12*time.Hour, 100*time.Minute),
	}
}
func (u *UserService) CreateUser(
	ctx context.Context,
	name string,
	email string,
	password string) (users.User, error) {
	if name == "" || email == "" || password == "" {
		return users.User{}, errors.New("fields must be filled")
	}
	user := users.User{
		FullName:     name,
		Email:        email,
		PasswordHash: password,
		CreatedAt:    time.Now(),
	}
	if err := u.repo.Create(ctx, &user); err != nil {
		return users.User{}, err
	}
	u.cache.Set(strconv.Itoa(user.ID), user, cache.DefaultExpiration)
	return user, nil
}
func (u UserService) GetAllService(ctx context.Context) (*[]users.User, error) {
	return u.repo.GetAll(ctx)
}
func (u UserService) GetByID(ctx context.Context, id int) (*users.User, error) {
	return u.repo.GetByID(ctx, id)
}
