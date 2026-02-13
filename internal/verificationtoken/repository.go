package verificationtoken

import (
	"context"
	"time"
)

type VerificationTokenRepository interface {
	Create(ctx context.Context, token VerificationToken) (VerificationToken, error)
	Get(ctx context.Context, id string) (VerificationToken, error)
	GetByToken(ctx context.Context, token string, tokenType string) (VerificationToken, error)
	Delete(ctx context.Context, id string) error
	DeleteByUser(ctx context.Context, userID string, tokenType string) error
	ListExpired(ctx context.Context, now time.Time) ([]VerificationToken, error)
}
