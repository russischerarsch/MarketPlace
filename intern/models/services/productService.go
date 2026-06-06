package services

import (
	"context"
	"fmt"
	"mini-ozon/intern/models/products"
	"mini-ozon/intern/repositories"
	"sync"
	"time"
)

const TimeLimit = 24 * time.Hour

type serviceCache struct {
	product     *products.Product
	expiresTime time.Time
}

type ProductService struct {
	repo  *repositories.ProductRepository
	cache map[int]serviceCache
	mtx   sync.Mutex
}

func (s *ProductService) SetProduct(id int, product *Product) {
	s.mtx.Lock()
	s.cache[id] = cacheItem{
		product:   product,
		expiresAt: time.Now().Add(defaultTTL), // ← текущее время + TTL
	}
	s.mtx.Unlock()
}
func (p *ProductService) CreateProduct(ctx context.Context, price int, desc string) (products.Product, error) {

	if price <= 0 {
		fmt.Println("price must be positive")
	}
	if desc == "" {
		fmt.Println("description required")
	}
	product := &products.Product{
		Price:       price,
		Description: desc,
		Created_at:  time.Now(),
	}
	p.mtx.Lock()
	p.cache[product.ID] = product
	p.mtx.Unlock()
	return *product, p.repo.Сreate(ctx, product)
}

func ttl(createdTime time.Time) {
	if time.Since(createdTime) == TimeLimit {

	}
}
