package users

import (
	"context"

	"gorm.io/gorm"
)

// UserService provides high level API to work with users
type UserService struct {
	client   *gorm.DB
	userRepo UserRepository
}

// NewUserService creates new UserService instance
func NewUserService(client *gorm.DB) *UserService {
	return &UserService{
		client:   client,
		userRepo: NewUserRepositoryPg(client),
	}
}

// GetUserByEmail returns user by provided Email
func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	return s.userRepo.GetByEmail(ctx, email)
}
