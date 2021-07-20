package users

// Role is a user role representation model
type Role struct {
	ID   uint64 `gorm:"primaryKey"`
	Name string
}

// TableName specifies table name for Role model
func (r *Role) TableName() string {
	return "auth.roles"
}
