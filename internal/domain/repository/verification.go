package repository

import (
	"context"
	"time"

	"github.com/scmbr/auth-service/internal/domain/entity"
)

type VerificationTokenRepository interface {
	Create(ctx context.Context, token entity.VerificationToken) (entity.VerificationToken, error)
	Get(ctx context.Context, id string) (entity.VerificationToken, error)
	GetByToken(ctx context.Context, token string, tokenType string) (entity.VerificationToken, error)
	Delete(ctx context.Context, id string) error
	DeleteByUser(ctx context.Context, userID string, tokenType string) error
	ListExpired(ctx context.Context, now time.Time) ([]entity.VerificationToken, error)
}
