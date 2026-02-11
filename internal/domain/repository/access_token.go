package repository

import (
	"context"

	"github.com/scmbr/auth-service/internal/domain/entity"
	"github.com/scmbr/auth-service/internal/usecase"
)

type AccessRepository interface {
	Create(ctx context.Context, accessToken entity.AccessToken) (entity.AccessToken, error)
	Get(ctx context.Context, id string) (entity.AccessToken, error)
	ListByFilter(ctx context.Context, filter usecase.AccessTokenFilter) ([]entity.AccessToken, error)
	Update(ctx context.Context, accessToken entity.AccessToken) (entity.AccessToken, error)
	Deactivate(ctx context.Context, id string) error
}
