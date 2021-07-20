package users

import "time"

// User is a user representation model
type User struct {
	ID        uint64 `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Email     string
	Pwd       string
	CreatedAt time.Time
	Roles     []*Role `gorm:"many2many:auth.users_roles;"`
}

// TableName specifies table name for User model
func (u *User) TableName() string {
	return "auth.users"
}
