package entity

import "time"

type Session struct {
	ID           string
	UserID       string
	ExpiresAt    time.Time
	IssuedAt     time.Time
	RefreshToken string
	Device       string
	IP           string
}
