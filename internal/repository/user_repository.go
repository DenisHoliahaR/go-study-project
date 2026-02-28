package repository

import (
	"context"

	"github.com/DenisHoliahaR/go-beautyhub/internal/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) (*domain.User, error)
	GetList(ctx context.Context) ([]*domain.User, error)
	GetByID(ctx context.Context, id int64) (*domain.User, error)
	Update(ctx context.Context, user *domain.User) (*domain.User, error)
	Delete(ctx context.Context, id int64) error
}