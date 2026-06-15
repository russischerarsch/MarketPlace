package services

import (
	"context"
	"errors"
	orderitem "mini-ozon/intern/models/orderItem"
	"mini-ozon/intern/repositories"
	"strconv"
	"time"

	"github.com/patrickmn/go-cache"
)

type OrderItemService struct {
	repo  *repositories.OrderItemRepository
	cache *cache.Cache
}

func CreateOrderItemService(repo *repositories.OrderItemRepository) *OrderItemService {
	return &OrderItemService{
		repo:  repo,
		cache: cache.New(12*time.Hour, 100*time.Minute),
	}
}
func (o OrderItemService) CreateOrderItem(
	ctx context.Context,
	OrderId int,
	ProductId int,
	Quantity int) (*orderitem.OrderItem, error) {
	if OrderId < 0 || Quantity < 0 || ProductId < 0 {
		return &orderitem.OrderItem{}, errors.New("must be positive")
	}
	tx, err := o.repo.BeginTx(ctx)
	if err != nil {
		return &orderitem.OrderItem{}, err
	}
	order := &orderitem.OrderItem{
		Order_id:   OrderId,
		Product_id: ProductId,
		Quantity:   Quantity,
	}

	if err := o.repo.CreateOrderItem(ctx, order, tx); err != nil {
		return &orderitem.OrderItem{}, err
	}
	o.cache.Set(strconv.Itoa(order.ID), order, cache.DefaultExpiration)
	return order, nil
}
func (o OrderItemService) GetAll(ctx context.Context) ([]orderitem.OrderItem, error) {
	return o.repo.GetAllOrderItem(ctx)
}
func (o OrderItemService) GetById(ctx context.Context, id int) (*orderitem.OrderItem, error) {
	if id < 0 {
		return &orderitem.OrderItem{}, errors.New("id must be positive")
	}
	if value, found := o.cache.Get(strconv.Itoa(id)); found {
		order, ok := value.(*orderitem.OrderItem)
		if !ok {
			return &orderitem.OrderItem{}, errors.New("failed to type Assertion")
		}
		return order, nil
	}
	return o.repo.GetByIdOrderItem(ctx, id)
}
