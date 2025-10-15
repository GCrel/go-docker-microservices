package postgres

import (
	"github.com/GCrel/go-microservices-docker/internal/users/core/domain"
	"gorm.io/gorm"
)

type PostgresRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) Save(user *domain.User) error {
	result := r.db.Create(user)
	return result.Error
}

func (r *PostgresRepository) FindByID(id string) (*domain.User, error) {
	var user domain.User
	result := r.db.First(&user, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}