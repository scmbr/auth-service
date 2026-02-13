package role

import (
	"context"

	"github.com/scmbr/auth-service/internal/scope"
)

type RoleRepository interface {
	Create(ctx context.Context, role Role) (Role, error)
	Get(ctx context.Context, id string) (Role, error)
	ListByFilter(ctx context.Context, filter RoleFilter) ([]Role, error)
	Update(ctx context.Context, role Role) (Role, error)
	Deactivate(ctx context.Context, id string) error

	AddScope(ctx context.Context, roleID, scopeID string) error
	RemoveScope(ctx context.Context, roleID, scopeID string) error
	ListScopes(ctx context.Context, roleID string) ([]scope.Scope, error)
	HasScope(ctx context.Context, roleID, scopeID string) (bool, error)
}
