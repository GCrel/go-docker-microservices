package ports

import "github.com/GCrel/go-microservices-docker/internal/products/core/domain"

type ProductService interface {
	CreateProduct(name, description string, price float64, sellerID string) (*domain.Product, error)
	GetProduct(id string) (*domain.Product, error)
}

type ProductRepository interface {
	Save(product *domain.Product) error
	GetByID(id string) (*domain.Product, error)
}