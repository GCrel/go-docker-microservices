package postgres

import (
	"github.com/GCrel/go-microservices-docker/internal/products/core/domain"
	"gorm.io/gorm"
)

type PostgresRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) Save(product *domain.Product) error {
	result := r.db.Create(product)
	return result.Error
}

func (r *PostgresRepository) GetByID(id string) (*domain.Product, error) {
	var product domain.Product
	result := r.db.First(&product, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}