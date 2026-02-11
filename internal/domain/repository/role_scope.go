package repository

import (
	"context"

	"github.com/scmbr/auth-service/internal/domain/entity"
)

type RoleScopeRepository interface {
	AddScope(ctx context.Context, roleID, permissionID string) error
	RemoveScope(ctx context.Context, roleID, permissionID string) error
	ListScopesByRole(ctx context.Context, roleID string) ([]entity.Scope, error)
	HasScope(ctx context.Context, roleID, permissionID string) (bool, error)
}
