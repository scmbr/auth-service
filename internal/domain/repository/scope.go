package repository

import (
	"context"

	"github.com/scmbr/auth-service/internal/domain/entity"
	"github.com/scmbr/auth-service/internal/usecase"
)

type ScopeRepository interface {
	Create(ctx context.Context, permission entity.Scope) (entity.Scope, error)
	Get(ctx context.Context, id string) (entity.Scope, error)
	ListByFilter(ctx context.Context, filter usecase.ScopeFilter) ([]entity.Scope, error)
	Update(ctx context.Context, permission entity.Scope) (entity.Scope, error)
	Deactivate(ctx context.Context, id string) error
}
