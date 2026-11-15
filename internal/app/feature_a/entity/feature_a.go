package entity

import "time"

// FeatureA TODO: ganti nama struct, tabel, dan field sesuai PRD.
type FeatureA struct {
	ID      uint `gorm:"primaryKey"`
	OwnerID uint `gorm:"index;not null"`
	// TODO: field sesuai PRD, contoh:
	// Title  string `gorm:"not null"`
	Status    string `gorm:"type:varchar(20);index;not null;default:draft"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Status enum. TODO: sesuaikan dengan alur status di PRD.
const (
	StatusDraft = "draft"
)
