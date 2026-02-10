package entity

import "time"

type VerificationToken struct {
	ID        string
	UserID    string
	Token     string
	ExpiresAt time.Time
	Type      string
}
