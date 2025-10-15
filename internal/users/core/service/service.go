package service

import (
	"time"

	"github.com/GCrel/go-microservices-docker/internal/users/core/domain"
	"github.com/GCrel/go-microservices-docker/internal/users/core/ports" 
	"github.com/google/uuid"
)

type UserServiceImpl struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		repo: repo,
	}
}

func (s *UserServiceImpl) CreateUser(name, email, password string) (*domain.User, error) {
	user := &domain.User{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
	}

	if err := s.repo.Save(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserServiceImpl) GetUserByID(id string) (*domain.User, error) {
	return s.repo.FindByID(id)
}