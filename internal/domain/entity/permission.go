package entity

type Permission struct {
	ID          string
	Name        string
	Description string
}
type RolePermission struct {
	RoleID       string
	PermissionID string
}
