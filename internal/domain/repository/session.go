package repository

import (
	"context"

	"github.com/scmbr/auth-service/internal/domain/entity"
	"github.com/scmbr/auth-service/internal/usecase"
)

type SessionRepository interface {
	Create(ctx context.Context, session entity.Session) (entity.Session, error)
	Get(ctx context.Context, id string) (entity.Session, error)
	ListByFilter(ctx context.Context, filter usecase.SessionFilter) ([]entity.Session, error)
	GetByDevice(ctx context.Context, userID *string, device string) (entity.Session, error)
	Update(ctx context.Context, session entity.Session) (entity.Session, error)
	Deactivate(ctx context.Context, id string) error
	DeactivateAllByUserID(ctx context.Context, userID string) error
}
