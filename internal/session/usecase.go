package session

type SessionFilter struct {
	UserID       *string
	ActiveOnly   bool
	DeactiveOnly bool
	Device       *string
	IP           *string
	Limit        int
	Offset       int
}
