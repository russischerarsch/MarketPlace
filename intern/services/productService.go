package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"mini-ozon/intern/models/products"

	"mini-ozon/intern/repositories"
	"time"

	"github.com/redis/go-redis/v9"
)

type ProductService struct {
	repo  *repositories.ProductRepository
	redis *redis.Client
}

func CreateProductService(repo *repositories.ProductRepository, redisClient *redis.Client) *ProductService {
	return &ProductService{
		repo:  repo,
		redis: redisClient,
	}
}
func (p *ProductService) CreateProduct(
	ctx context.Context,
	price int,
	desc string,
	name string) (products.Product, error) {

	if price <= 0 {
		return products.Product{}, errors.New("price must be positive")
	}
	if desc == "" || name == "" {
		return products.Product{}, errors.New("fields must be filled")
	}
	product := &products.Product{
		Name:        name,
		Price:       price,
		Description: desc,
		Created_at:  time.Now(),
	}
	if err := p.repo.Create(ctx, product); err != nil {
		return products.Product{}, err
	}
	key := fmt.Sprintf("product:%d", product.ID)
	data, _ := json.Marshal(product)
	p.redis.Set(ctx, key, data, 1*time.Hour)
	return *product, nil
}
func (p *ProductService) GetAll(ctx context.Context) ([]products.Product, error) {
	return p.repo.GetAll(ctx)
}

func (p *ProductService) GetByID(ctx context.Context, id int) (products.Product, error) {
	if id < 0 {
		return products.Product{}, errors.New("id must be positive")
	}
	key := fmt.Sprintf("product:%d", id)
	val, err := p.redis.Get(ctx, key).Result()
	if err == nil {
		var product products.Product
		json.Unmarshal([]byte(val), &product)
		return product, nil
	}
	product, err := p.repo.GetByID(ctx, id)
	if err != nil {
		return products.Product{}, err
	}
	data, _ := json.Marshal(product)
	p.redis.Set(ctx, key, data, 1*time.Hour)
	return p.repo.GetByID(ctx, id)
}
func (p ProductService) DeleteByID(ctx context.Context, id int) error {
	if id < 0 {
		return errors.New("id must be positive")
	}
	err := p.repo.DeleteByID(ctx, id)
	key := fmt.Sprintf("product:%d", id)
	p.redis.Del(ctx, key)
	return err
}
