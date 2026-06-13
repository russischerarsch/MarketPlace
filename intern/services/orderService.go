package services

import (
	"context"
	"errors"
	orderitem "mini-ozon/intern/models/orderItem"
	"mini-ozon/intern/models/orders"
	"mini-ozon/intern/repositories"
	"strconv"
	"time"

	"github.com/patrickmn/go-cache"
)

type OrderService struct {
	orderrepo   *repositories.OrderRepository
	productrepo *repositories.ProductRepository
	cache       *cache.Cache
}

func CreateOrderService(repo *repositories.OrderRepository) *OrderService {
	return &OrderService{
		orderrepo: repo,
		cache:     cache.New(12*time.Hour, 100*time.Minute),
	}
}
func (p *OrderService) CreateOrder(
	ctx context.Context,
	userID int,
	orderItems []orderitem.OrderItemRequest,
) (orders.Order, error) {
	total := 0
	for _, item := range orderItems {
		product, _ := p.productrepo.GetByID(ctx, item.ProductID)
		total += product.Price
	}
	order := orders.Order{
		UserID:    userID,
		Total:     total,
		Status:    orders.Pending,
		CreatedAt: time.Now(),
	}
	if err := p.orderrepo.Create(ctx, &order); err != nil {
		return orders.Order{}, err
	}
	p.cache.Set(strconv.Itoa(order.ID), order, cache.DefaultExpiration)
	return order, nil
}
func (p *OrderService) GetAll(ctx context.Context) ([]orders.Order, error) {
	return p.orderrepo.GetAll(ctx)

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
	return p.orderrepo.GetByID(ctx, id)
}
func (p OrderService) GetByUserID(ctx context.Context, id int) ([]orderitem.OrderWithItem, error) {
	var result []orderitem.OrderWithItem
	orders, err := p.orderrepo.GetAllByUserId(ctx, id)
	if err != nil {
		return nil, err
	}
	for _, order := range orders {
		items, err := p.orderrepo.GetOrderItems(ctx, order.ID)
		if err != nil {
			return nil, err
		}
		o := orderitem.OrderWithItem{
			Order:     order,
			OrderItem: items,
		}
		result = append(result, o)

	}
	return result, nil
}
