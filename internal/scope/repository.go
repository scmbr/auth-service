package scope

import (
	"context"
)

type ScopeRepository interface {
	Create(ctx context.Context, permission Scope) (Scope, error)
	Get(ctx context.Context, id string) (Scope, error)
	ListByFilter(ctx context.Context, filter ScopeFilter) ([]Scope, error)
	Update(ctx context.Context, permission Scope) (Scope, error)
	Deactivate(ctx context.Context, id string) error
}
