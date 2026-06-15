// internal/services/product_service_test.go
package services

import (
	"context"
	"errors"
	"mini-ozon/intern/models/products"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// ========== ИНТЕРФЕЙС ==========
type ProductRepository interface {
	Create(ctx context.Context, p *products.Product) error
	GetByID(ctx context.Context, id int) (*products.Product, error)
}

// ========== СЕРВИС ==========
type ProductServiceTest struct {
	repo ProductRepository
}

func NewProductServiceTest(repo ProductRepository) *ProductServiceTest {
	return &ProductServiceTest{repo: repo}
}

func (p *ProductServiceTest) CreateProd(ctx context.Context, price int, desc string, name string) (products.Product, error) {
	if name == "" {
		return products.Product{}, errors.New("name is required")
	}
	if price <= 0 {
		return products.Product{}, errors.New("price must be positive")
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

	return *product, nil
}

func (p *ProductServiceTest) GetProduct(ctx context.Context, id int) (*products.Product, error) {
	if id <= 0 {
		return &products.Product{}, errors.New("invalid id")
	}
	return p.repo.GetByID(ctx, id)
}

// ========== МОК ==========
type MockProductRepo struct {
	mock.Mock
}

func (m *MockProductRepo) Create(ctx context.Context, p *products.Product) error {
	args := m.Called(ctx, p)
	return args.Error(0)
}

func (m *MockProductRepo) GetByID(ctx context.Context, id int) (*products.Product, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*products.Product), args.Error(1)
}

// ========== ТЕСТЫ ==========
func TestCreateProduct_Success(t *testing.T) {
	mockRepo := new(MockProductRepo)
	service := NewProductServiceTest(mockRepo)

	mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*products.Product")).Return(nil)

	product, err := service.CreateProduct(context.Background(), 100, "test description", "Test")

	assert.NoError(t, err)
	assert.Equal(t, "Test", product.Name)
	assert.Equal(t, 100, product.Price)
	assert.Equal(t, "test description", product.Description)
	mockRepo.AssertExpectations(t)
}

func TestCreateProduct_EmptyName(t *testing.T) {
	mockRepo := new(MockProductRepo)
	service := NewProductServiceTest(mockRepo)

	_, err := service.CreateProduct(context.Background(), 100, "desc", "")

	assert.Error(t, err)
	assert.Equal(t, "name is required", err.Error())
}

func TestCreateProduct_InvalidPrice(t *testing.T) {
	mockRepo := new(MockProductRepo)
	service := NewProductServiceTest(mockRepo)

	_, err := service.CreateProduct(context.Background(), -10, "desc", "Test")

	assert.Error(t, err)
	assert.Equal(t, "price must be positive", err.Error())
}

func TestCreateProduct_RepoError(t *testing.T) {
	mockRepo := new(MockProductRepo)
	service := NewProductServiceTest(mockRepo)

	mockRepo.On("Create", mock.Anything, mock.Anything).Return(errors.New("db error"))

	_, err := service.CreateProduct(context.Background(), 100, "desc", "Test")

	assert.Error(t, err)
	assert.Equal(t, "db error", err.Error())
}

func TestGetProduct_Success(t *testing.T) {
	mockRepo := new(MockProductRepo)
	service := NewProductServiceTest(mockRepo)

	expected := &products.Product{ID: 1, Name: "Test", Price: 100}
	mockRepo.On("GetByID", mock.Anything, 1).Return(expected, nil)

	product, err := service.GetProduct(context.Background(), 1)

	assert.NoError(t, err)
	assert.Equal(t, "Test", product.Name)
	mockRepo.AssertExpectations(t)
}

func TestGetProduct_InvalidID(t *testing.T) {
	mockRepo := new(MockProductRepo)
	service := NewProductServiceTest(mockRepo)

	_, err := service.GetProduct(context.Background(), 0)

	assert.Error(t, err)
	assert.Equal(t, "invalid id", err.Error())
}

func TestGetProduct_NotFound(t *testing.T) {
	mockRepo := new(MockProductRepo)
	service := NewProductServiceTest(mockRepo)

	mockRepo.On("GetByID", mock.Anything, 999).Return(nil, errors.New("not found"))

	_, err := service.GetProduct(context.Background(), 999)

	assert.Error(t, err)
}
