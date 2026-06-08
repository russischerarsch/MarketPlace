package services

import (
	"context"
	"errors"
	"mini-ozon/intern/models/orders"
	"mini-ozon/intern/repositories"
	"strconv"
	"time"

	"github.com/patrickmn/go-cache"
)

type OrderService struct {
	repo  *repositories.OrderRepository
	cache *cache.Cache
}

func CreateOrderService(repo *repositories.OrderRepository) *OrderService {
	return &OrderService{
		repo:  repo,
		cache: cache.New(12*time.Hour, 100*time.Minute),
	}
}
func (p *OrderService) CreateOrder(
	ctx context.Context,
	userID int,
) (orders.Order, error) {
	order := orders.Order{
		UserID:    userID,
		Status:    orders.Pending,
		CreatedAt: time.Now(),
	}
	if err := p.repo.Create(ctx, userID, &order); err != nil {
		return orders.Order{}, err
	}
	p.cache.Set(strconv.Itoa(order.ID), order, cache.DefaultExpiration)
	return order, nil
}
func (p *OrderService) GetAll(ctx context.Context) ([]orders.Order, error) {
	return p.repo.GetAll(ctx)

}

func (p *OrderService) GetByID(ctx context.Context, id int) (orders.Order, error) {
	if id < 0 {
		return orders.Order{}, errors.New("id must be positive")
	}
	if value, found := p.cache.Get(strconv.Itoa(id)); found {
		order, ok := value.(*orders.Order)
		if !ok {
			return orders.Order{}, errors.New("failed to type Assertion")
		}
		return *order, nil
	}
	return p.repo.GetByID(ctx, id)
}
