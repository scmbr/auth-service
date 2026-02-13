package user

import (
	"context"

	"github.com/scmbr/auth-service/internal/role"
	"github.com/scmbr/auth-service/internal/session"
)

type UserUsecase struct {
	userRepo    UserRepository
	sessionRepo session.SessionRepository
	roleRepo    role.RoleRepository
}

func NewUserUsecase(userRepo UserRepository, sessionRepo session.SessionRepository, roleRepo role.RoleRepository) *UserUsecase {
	return &UserUsecase{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
		roleRepo:    roleRepo,
	}
}

type LoginInput struct {
	phoneNumber *string
	email       *string
	password    string
}
type RegisterInput struct {
	phoneNumber *string
	email       *string
	password    string
}
type ChangePasswordInput struct {
	phoneNumber *string
	email       *string
	oldPassword string
	newPassword string
}
type CreateInput struct {
	email    string
	phone    string
	password string
}
type UserFilter struct {
	ActiveOnly   bool
	DeactiveOnly bool
	Offset       int
	Limit        int
}

func (uc *UserUsecase) Login(ctx context.Context, in LoginInput) (*User, error) {
	return nil, nil
}

func (uc *UserUsecase) Register(ctx context.Context, in RegisterInput) (*User, error) {
	return nil, nil

}
func (uc *UserUsecase) ChangePassword(ctx context.Context, in ChangePasswordInput) error {
	return nil

}
func (uc *UserUsecase) Create(ctx context.Context, in CreateInput) (*User, error) {
	return nil, nil
}
func (uc *UserUsecase) GetById(ctx context.Context, id string) (*User, error) {
	return nil, nil
}
func (uc *UserUsecase) Activate(ctx context.Context, id string) error {
	return nil
}
func (uc *UserUsecase) Deactivate(ctx context.Context, id string) error {
	return nil
}
func (uc *UserUsecase) AddRole(ctx context.Context, id string, roleID string) error {
	return nil
}
func (uc *UserUsecase) RemoveRole(ctx context.Context, id string, roleID string) error {
	return nil
}
func (uc *UserUsecase) HasRole(ctx context.Context, id string, roleID string) (*bool, error) {
	return nil, nil
}
func (uc *UserUsecase) ListRoles(ctx context.Context, id string) ([]*role.Role, error) {
	return nil, nil
}
