package entity

type Role struct {
	ID          string
	Name        string
	Description string
}
type UserRole struct {
	UserID string
	RoleID string
}
