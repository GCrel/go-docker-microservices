package ports

import "github.com/GCrel/go-microservices-docker/internal/users/core/domain"

type UserService interface {
	CreateUser(name string, email string, password string) (*domain.User, error)
	GetUserByID(id string) (*domain.User, error)
}

type UserRepository interface {
	Save(user *domain.User) error
	FindByID(id string) (*domain.User, error)
}