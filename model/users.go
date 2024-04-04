package model

type UserRole int

const (
	Admin UserRole = iota
	Users
)

type User struct {
	Model
	Username string   `gorm:"not null" json:"username"`
	Name     string   `gorm:"not null" json:"name"`
	Password string   `gorm:"not null" json:"password"`
	Role     UserRole `gorm:"not null" json:"role"`
}
