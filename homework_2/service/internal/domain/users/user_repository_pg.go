package users

import (
	"context"

	"gorm.io/gorm"
)

// UserRepositoryPg is a repository for User model over PostgeSQL
// Implements UserRepository
type UserRepositoryPg struct {
	client *gorm.DB
}

// NewUserRepositoryPg creates new NewUserRepositoryPg instance
func NewUserRepositoryPg(client *gorm.DB) *UserRepositoryPg {
	return &UserRepositoryPg{
		client: client,
	}
}

// GetByID returns user by provided ID
func (r *UserRepositoryPg) GetByID(ctx context.Context, id uint64) (*User, error) {
	var user User
	if err := r.client.WithContext(ctx).First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// GetByEmail returns user by provided Email
func (r *UserRepositoryPg) GetByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	if err := r.client.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
