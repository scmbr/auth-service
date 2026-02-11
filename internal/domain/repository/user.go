package repository

import (
	"context"

	"github.com/scmbr/auth-service/internal/domain/entity"
	"github.com/scmbr/auth-service/internal/usecase"
)

type UserRepository interface {
	Create(ctx context.Context, user entity.User) (entity.User, error)
	ListByFilter(ctx context.Context, filter usecase.UserFilter) ([]entity.User, error)
	Get(ctx context.Context, id string) (entity.User, error)
	GetByEmail(ctx context.Context, email string) (entity.User, error)
	GetByUsername(ctx context.Context, username string) (entity.User, error)
	GetByIDs(ctx context.Context, ids []string) ([]entity.User, error)
	Update(ctx context.Context, user entity.User) (entity.User, error)
	Deactivate(ctx context.Context, id string) error
	Activate(ctx context.Context, id string) error
	IsActive(ctx context.Context, id string) (bool, error)
}
