package services

import (
	"context"
	"errors"
	"fmt"
	"mini-ozon/intern/models/products"
	"mini-ozon/intern/repositories"
	"strconv"
	"time"

	"github.com/patrickmn/go-cache"
)

const TimeLimit = 48 * time.Hour

type ProductService struct {
	repo  *repositories.ProductRepository
	cache *cache.Cache
}

func Create(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{
		repo:  repo,
		cache: cache.New(12*time.Hour, 100*time.Minute),
	}
}
func (p *ProductService) CreateProduct(
	ctx context.Context,
	price int,
	desc string,
	name string) (products.Product, error) {

	if price <= 0 {
		fmt.Println("price must be positive")
	}
	if desc == "" || name == "" {
		fmt.Println("field must be filled")
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
	p.cache.Set(strconv.Itoa(product.ID), product, cache.DefaultExpiration)
	return *product, nil
}
func (p *ProductService) GetAll(ctx context.Context) ([]products.Product, error) {
	return p.repo.GetAll(ctx)
}

func (p *ProductService) GetByID(ctx context.Context, id int) (products.Product, error) {
	if id < 0 {
		return products.Product{}, errors.New("id must be positive")
	}
	if value, found := p.cache.Get(strconv.Itoa(id)); found {
		product, ok := value.(*products.Product)
		if !ok {
			return products.Product{}, errors.New("failed to type Assertion")
		}
		return *product, nil
	}
	return p.repo.GetByID(ctx, id)
}
