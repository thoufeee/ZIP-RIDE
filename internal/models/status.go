package models

import (
	"github.com/google/uuid"
	"time"
)

type Status struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	DriverID  uuid.UUID `gorm:"type:uuid;uniqueIndex" json:"driver_id"`
	Available bool      `json:"available"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	UpdatedAt time.Time `json:"updated_at"`
}
