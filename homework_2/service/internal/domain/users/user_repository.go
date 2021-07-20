package users

import (
	"context"
)

// UserRepository is a repository abstraction for User model
type UserRepository interface {
	// GetByID returns user by provided ID
	GetByID(ctx context.Context, id uint64) (*User, error)

	// GetByEmail returns user by provided Email
	GetByEmail(ctx context.Context, email string) (*User, error)
}
