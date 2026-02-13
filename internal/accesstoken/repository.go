package accesstoken

import (
	"context"
)

type AccessRepository interface {
	Create(ctx context.Context, accessToken AccessToken) (AccessToken, error)
	Get(ctx context.Context, id string) (AccessToken, error)
	ListByFilter(ctx context.Context, filter AccessTokenFilter) ([]AccessToken, error)
	Update(ctx context.Context, accessToken AccessToken) (AccessToken, error)
	Deactivate(ctx context.Context, id string) error
}
