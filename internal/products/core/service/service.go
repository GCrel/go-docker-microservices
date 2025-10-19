package service

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/GCrel/go-microservices-docker/internal/products/core/domain"
	"github.com/GCrel/go-microservices-docker/internal/products/core/ports"
	"github.com/google/uuid"
)

type ProductServiceImpl struct {
	repo        ports.ProductRepository
	usersApiURL string
}

func NewProductService(repo ports.ProductRepository, usersApiURL string) *ProductServiceImpl {
	return &ProductServiceImpl{
		repo:        repo,
		usersApiURL: usersApiURL,
	}
}

func (s *ProductServiceImpl) CreateProduct(name, description string, price float64, sellerID string) (*domain.Product, error) {

	userURL := fmt.Sprintf("%s/users/%s", s.usersApiURL, sellerID)

	fmt.Println(userURL)
	resp, err := http.Get(userURL)
	if err != nil {
		return nil, fmt.Errorf("error al contactar el servicio de usuarios: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("el vendedor con el ID proporcionado no existe")
	}

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
