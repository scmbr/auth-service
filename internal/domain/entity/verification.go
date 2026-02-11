package entity

import "time"

type VerificationTokenType string

const (
	TokenTypeEmailVerification VerificationTokenType = "email_verification"
	TokenTypePasswordReset     VerificationTokenType = "password_reset"
	TokenType2FA               VerificationTokenType = "2fa"
)

type VerificationToken struct {
	ID        string
	UserID    string
	Token     string
	ExpiresAt time.Time
	Type      VerificationTokenType
}
