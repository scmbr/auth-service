package user

import (
	"github.com/scmbr/auth-service/internal/role"
	"github.com/scmbr/auth-service/internal/session"
)

type UserUsecase struct {
	userRepo    UserRepository
	sessionRepo session.SessionRepository
	roleRepo    role.RoleRepository
}

type UserFilter struct {
}
