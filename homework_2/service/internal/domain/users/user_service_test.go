// +build integration
package users_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"go-postgres/internal/domain/users"

	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	adminRole = &users.Role{
		ID:   1,
		Name: "admin",
	}

	preparedData = []*users.User{
		{
			ID:        1,
			FirstName: "User",
			LastName:  "One",
			Email:     "test@usermail.org",
			Pwd:       "",
			CreatedAt: time.Now(),
			Roles: []*users.Role{
				adminRole,
			},
		},
	}
)

// TestGetUserByEmail performs GetUserByEmail method integration test
func TestGetUserByEmail(t *testing.T) {
	// test cases
	tests := []struct {
		name    string
		ctx     context.Context
		email   string
		prepare func(db *gorm.DB)
		check   func(*testing.T, *users.User, error)
	}{
		{
			name:  "success",
			ctx:   context.Background(),
			email: "test@usermail.org",

			// preparing data for test
			prepare: func(db *gorm.DB) {
				db.WithContext(context.Background()).Create(adminRole)
				db.WithContext(context.Background()).Create(preparedData)
			},

			// check results
			check: func(t *testing.T, user *users.User, err error) {
				require.NoError(t, err)
				require.NotNil(t, user)
			},
		},

		{
			name:    "no user found",
			ctx:     context.Background(),
			email:   "not_exists@usermail.org",
			prepare: func(db *gorm.DB) {},

			// check results
			check: func(t *testing.T, user *users.User, err error) {
				require.NoError(t, err)
				require.Nil(t, user)
			},
		},
	}

	// connect to database and create UserService
	db := connect()
	service := users.NewUserService(db)

	// do tests
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare(db)
			user, err := service.GetUserByEmail(tt.ctx, tt.email)
			tt.check(t, user, err)
		})
	}
}

// connect makes connection with postgres
func connect() *gorm.DB {
	dbhost := os.Getenv("DATABASE_HOST")
	dbport := os.Getenv("DATABASE_PORT")
	dbuser := os.Getenv("DATABASE_USER")
	dbpwd := os.Getenv("DATABASE_PASSWORD")
	dbname := os.Getenv("DATABASE_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbhost, dbport, dbuser, dbpwd, dbname)
	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)

	if err != nil {
		panic(err)
	}

	return db
}
