package repository

import (
	"context"

	"github.com/scmbr/auth-service/internal/domain/entity"
)

type UserRoleRepository interface {
	AddRole(ctx context.Context, userID, roleID string) error
	RemoveRole(ctx context.Context, userID, roleID string) error
	ListRolesByUser(ctx context.Context, userID string) ([]entity.Role, error)
	HasRole(ctx context.Context, userID, roleID string) (bool, error)
}
