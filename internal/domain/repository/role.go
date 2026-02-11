package repository

import (
	"context"

	"github.com/scmbr/auth-service/internal/domain/entity"
	"github.com/scmbr/auth-service/internal/usecase"
)

type Role interface {
	Create(ctx context.Context, role entity.Role) (entity.Role, error)
	Get(ctx context.Context, id string) (entity.Role, error)
	ListByFilter(ctx context.Context, filter usecase.RoleFilter) ([]entity.Role, error)
	Update(ctx context.Context, role entity.Role) (entity.Role, error)
	Deactivate(ctx context.Context, id string) error
}
