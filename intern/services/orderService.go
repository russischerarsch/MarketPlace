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
	orderrepo     *repositories.OrderRepository
	productrepo   *repositories.ProductRepository
	orderitemrepo *repositories.OrderItemRepository
	cache         *cache.Cache
}

func CreateOrderService(repo *repositories.OrderRepository,
	prodrepo *repositories.ProductRepository,
	orderitemrepo *repositories.OrderItemRepository) *OrderService {
	return &OrderService{
		orderrepo:     repo,
		productrepo:   prodrepo,
		orderitemrepo: orderitemrepo,
		cache:         cache.New(12*time.Hour, 100*time.Minute),
	}
}
func (p *OrderService) CreateOrder(
	ctx context.Context,
	userID int,
	orderItems []orderitem.OrderItemRequest,
) (orders.Order, error) {
	tx, err := p.orderrepo.BeginTx(ctx)
	if err != nil {
		return orders.Order{}, err
	}
	defer tx.Rollback(ctx)
	total := 0
	for _, item := range orderItems {
		if item.Quantity <= 0 {
			return orders.Order{}, errors.New("quantity must be positive")
		}
		product, err := p.productrepo.GetByID(ctx, item.ProductID)
		if err != nil {
			return orders.Order{}, err
		}
		total += product.Price * item.Quantity
	}
	order := orders.Order{
		UserID:    userID,
		Total:     total,
		Status:    orders.Pending,
		CreatedAt: time.Now(),
	}
	if err := p.orderrepo.Create(ctx, &order, tx); err != nil {
		return orders.Order{}, err
	}
	for _, item := range orderItems {
		itemSt := orderitem.OrderItem{
			Order_id:   order.ID,
			Product_id: item.ProductID,
			Quantity:   item.Quantity,
		}
		if err := p.orderitemrepo.CreateOrderItem(ctx, &itemSt, tx); err != nil {
			return orders.Order{}, err
		}
	}
	if err := tx.Commit(ctx); err != nil {
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
