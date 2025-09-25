package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Driver struct {
	ID            uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name          string         `json:"name" binding:"required"`
	Email         string         `gorm:"uniqueIndex" json:"email"`
	Phone         string         `gorm:"uniqueIndex" json:"phone" binding:"required"`
	Password      string         `json:"-"`
	LicenseNo     string         `gorm:"uniqueIndex" json:"license_no"`
	VehicleID     uint           `json:"vehicle_id"`
	GoogleID      string         `gorm:"uniqueIndex" json:"google_id,omitempty"`
	PhoneVerified bool           `gorm:"default:false" json:"phone_verified"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
