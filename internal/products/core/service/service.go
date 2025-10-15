package service

import (
	"time"

	"github.com/GCrel/go-microservices-docker/internal/products/core/domain"
	"github.com/GCrel/go-microservices-docker/internal/products/core/ports"
	"github.com/google/uuid"
)

type ProductServiceImpl struct {
	repo ports.ProductRepository
}

func NewProductService(repo ports.ProductRepository) *ProductServiceImpl {
	return &ProductServiceImpl{
		repo: repo,
	}
}

func (s *ProductServiceImpl) CreateProduct(name, description string, price float64, sellerID string) (*domain.Product, error) {
	product := &domain.Product{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		Price:       price,
		SellerID:    sellerID,
		CreatedAt:   time.Now(),
	}

	if err := s.repo.Save(product); err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductServiceImpl) GetProduct(id string) (*domain.Product, error) {
	return s.repo.GetByID(id)
}