package session

import (
	"context"
)

type SessionRepository interface {
	Create(ctx context.Context, session Session) (Session, error)
	Get(ctx context.Context, id string) (Session, error)
	ListByFilter(ctx context.Context, filter SessionFilter) ([]Session, error)
	GetByDevice(ctx context.Context, userID *string, device string) (Session, error)
	Update(ctx context.Context, session Session) (Session, error)
	Deactivate(ctx context.Context, id string) error
	DeactivateAllByUserID(ctx context.Context, userID string) error
}
