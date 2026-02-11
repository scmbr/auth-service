package session

import "time"

type Session struct {
	ID            string
	UserID        *string
	ExpiresAt     time.Time
	IssuedAt      time.Time
	DeactivatedAt *time.Time
	RefreshToken  string
	DeviceID      string
	UserAgent     string
	IP            string
}
