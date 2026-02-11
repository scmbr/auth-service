package entity

type Scope struct {
	ID          string
	Name        string
	Description string
}
type RoleScope struct {
	RoleID  string
	ScopeID string
}
