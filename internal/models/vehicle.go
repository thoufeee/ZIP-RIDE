package models

import (
	"github.com/google/uuid"
	"time"
)

type Vehicle struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Make        string    `json:"make"`
	Model       string    `json:"model"`
	Year        int       `json:"year"`
	PlateNumber string    `gorm:"uniqueIndex" json:"plate_number"`
	Color       string    `json:"color"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	Driver Driver `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
