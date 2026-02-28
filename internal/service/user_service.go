package services

import (
	"context"
	"fmt"

	"github.com/DenisHoliahaR/go-beautyhub/internal/domain"
	"github.com/DenisHoliahaR/go-beautyhub/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) *UserService {
	return &UserService{
		repo: r,
	}
}

func (s *UserService) CreateUser(ctx context.Context, user *domain.User, password string) (*domain.User, error) {
	if len(password) < 8 {
		return nil, fmt.Errorf("weak password")
	}

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, fmt.Errorf("hash password: %w", err)
	}

	user.PasswordHash = string(hash)

	createdUser, err := s.repo.Create(ctx, user); 
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (s *UserService) GetUserById(ctx context.Context, id int64) (*domain.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *UserService) GetUserList(ctx context.Context) ([]*domain.User, error) {
	return s.repo.GetList(ctx)
}

func (s *UserService) UpdateUser(ctx context.Context, user *domain.User, password string) (*domain.User, error) {
	if password != "" {
		if len(password) < 8 {
			return nil, fmt.Errorf("weak password")
		}

		hash, err := bcrypt.GenerateFromPassword(
			[]byte(password),
			bcrypt.DefaultCost,
		)
		if err != nil {
			return nil, fmt.Errorf("hash password: %w", err)
		}
	
		user.PasswordHash = string(hash)
	}

	updatedUser, err := s.repo.Update(ctx, user); 
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
