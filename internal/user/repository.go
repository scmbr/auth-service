package user

import (
	"context"

	"github.com/scmbr/auth-service/internal/role"
)

type UserRepository interface {
	Create(ctx context.Context, user User) (User, error)
	ListByFilter(ctx context.Context, filter UserFilter) ([]User, error)
	Get(ctx context.Context, id string) (User, error)
	GetByEmail(ctx context.Context, email string) (User, error)
	GetByUsername(ctx context.Context, username string) (User, error)
	GetByIDs(ctx context.Context, ids []string) ([]User, error)
	Update(ctx context.Context, user User) (User, error)
	Deactivate(ctx context.Context, id string) error
	Activate(ctx context.Context, id string) error
	IsActive(ctx context.Context, id string) (bool, error)

	AddRole(ctx context.Context, userID, roleID string) error
	RemoveRole(ctx context.Context, userID, roleID string) error
	ListRoles(ctx context.Context, userID string) ([]role.Role, error)
	HasRole(ctx context.Context, userID, roleID string) (bool, error)
}
