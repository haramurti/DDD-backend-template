package entity

import "time"

const (
	RoleMember  = "member"
	RoleManager = "manager" // TODO: sesuaikan nama role dengan PRD
)

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Name         string `gorm:"not null"`
	Email        string `gorm:"uniqueIndex;not null"`
	PasswordHash string `gorm:"not null"`
	Role         string `gorm:"type:varchar(20);not null;default:member"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
