package accesstoken

import "time"

type AccessToken struct {
	ID        string
	UserID    string
	RefreshID string
	CreatedAt time.Time
	RevokedAt *time.Time
	ExpiresAt time.Time
	RevokedBy *string
}
